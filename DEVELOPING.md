# Project Development

For general information about contributing changes, see the
[Contributor Guidelines](https://github.com/ditdotdev/.github/blob/master/CONTRIBUTING.md).

## How it Works

The Go client is generated from `dit.yml` by [openapi-generator-cli](https://openapi-generator.tech/),
pinned to **v7.22.0** in the `generate` script. The generated files are committed to the repository so
they can be imported as a Go module without consumers needing the generator tooling.

To regenerate after spec changes:

1. Update `dit.yml` with the new specification.
2. Run `./generate`. This requires Docker (the script runs the openapi-generator container).
3. If the new spec removes models or operations, manually `git rm` the corresponding `model_*.go` /
   `api_*.go` files — the generator does not delete stale files on its own.
4. Run `go build ./...` and `go test ./...` to verify the regeneration is clean.

### Templates

We **do not** maintain local mustache template overrides. The stock upstream Go templates emit
idiomatic v7 code (fluent request builders, `io.ReadAll`, full marshallers, getter/setter pairs).
Customisation that needs to differ from stock is passed via `--additional-properties` in the
`generate` script — currently `packageName`, `withGoCodegenComment`, and `apiNameSuffix=Api`.

If a future generator upgrade emits something unacceptable, the right fix is to either (a) negotiate
upstream, or (b) add a single targeted override under a fresh `templates/` directory. Avoid wholesale
template copies — they go stale and turn the next generator upgrade into archaeology.

## Production-Safe Defaults

The stock openapi-generator output ships with two latent issues that consumers should be aware of.
This package adds hand-written helpers in [`safe_defaults.go`](safe_defaults.go) that work around
them without re-introducing local mustache template overrides.

### HTTP client timeout

`NewAPIClient` assigns `http.DefaultClient` when `cfg.HTTPClient` is nil. That is a process-wide
singleton with no `Timeout`; a hung server hangs the calling goroutine indefinitely, and any
mutation of `http.DefaultClient` leaks into every other package in the process.

Use `NewAPIClientWithDefaults` to get a non-shared `*http.Client` with a generous default
timeout, or pass your own `cfg.HTTPClient` for finer control:

```go
cfg := ditclient.NewConfiguration()
// Use the safer constructor — fresh *http.Client, 5 minute timeout.
client := ditclient.NewAPIClientWithDefaults(cfg)

// Or supply your own client.
cfg.HTTPClient = &http.Client{Timeout: 10 * time.Second}
client = ditclient.NewAPIClient(cfg)
```

See [issue #32](https://github.com/ditdotdev/dit-client-go/issues/32).

### Default headers and concurrency

`Configuration.DefaultHeader` is an unguarded `map[string]string`. The generated `AddDefaultHeader`
writes it without a lock, and the generated request-prepare path iterates it on every request.
Mutating headers from a goroutine while requests are in flight will trip Go's race detector and
may panic with "concurrent map iteration and map write".

The robust pattern is to **set headers once at init**, before constructing the API client:

```go
cfg := ditclient.NewConfiguration()
ditclient.SetDefaultHeader(cfg, "X-Tenant", tenantID)
client := ditclient.NewAPIClientWithDefaults(cfg)
```

If you must mutate headers later (e.g. a token-refresh goroutine), use `SetDefaultHeader` and
`DefaultHeaders` from this package — they share a global RWMutex so writes don't race with each
other. **They do not fully protect against the generated iteration**; the upstream fix is tracked
in [issue #31](https://github.com/ditdotdev/dit-client-go/issues/31).

## Using the Generated Client

The v7 generator emits a fluent request-builder API. Each operation returns a request type that
you can decorate with optional parameters before calling `.Execute()`:

```go
cfg := ditclient.NewConfiguration()
cfg.Servers = ditclient.ServerConfigurations{
    {URL: "http://localhost:5001"},
}
client := ditclient.NewAPIClient(cfg)

// GET with required path params only
commit, resp, err := client.CommitsApi.
    GetCommit(ctx, "myrepo", "abc123").
    Execute()

// POST with a required request body
created, resp, err := client.CommitsApi.
    CreateCommit(ctx, "myrepo").
    Commit(ditclient.Commit{Id: "new", Properties: map[string]any{}}).
    Execute()

// GET with optional query parameters
commits, resp, err := client.CommitsApi.
    ListCommits(ctx, "myrepo").
    Tag([]string{"production"}).
    Execute()
```

Errors are returned as `*GenericOpenAPIError`. When the server responds with an `application/json`
error body that matches the `ApiError` schema, the parsed model is available via `err.Model()`.

## Building

```bash
go build ./...
go test ./...
```

## Testing

Hand-written integration tests live alongside the generated code as `*_test.go` files (e.g.
`api_commits_test.go`). They spin up an `httptest.NewServer` so we can exercise the real request
path, body marshalling, and error-decoding behaviour without depending on a live `dit-server`.

The openapi-generator also emits its own auto-stub test files at `test/api_*_test.go`. We ignore
those via `.openapi-generator-ignore` — they wrap every operation in `t.Skip()` and only assert
`StatusCode == 200`, so they add no real coverage.

## Releasing

The module is released by tagging the commit. CI publishes a GitHub Release via release-drafter
and clients pick up the new version through `go get`.

```bash
git tag v1.0.0
git push origin v1.0.0
```

Consumer repos (`dit`, `dit-server`, `dit-docker-proxy`) bump
`github.com/ditdotdev/dit-client-go` in their `go.mod` after a release.
