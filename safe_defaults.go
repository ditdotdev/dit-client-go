package datadatdatclient

import (
	"net/http"
	"sync"
	"time"
)

// This file holds hand-written safety helpers that complement the
// auto-generated client. They exist because two latent issues in the
// upstream openapi-generator-cli output (v7.22.0) cannot be fixed
// without re-introducing local mustache template overrides:
//
//   - Configuration.DefaultHeader is an unguarded map[string]string,
//     mutated by AddDefaultHeader and iterated by prepareRequest on every
//     request. Concurrent writes during in-flight requests will trip the
//     Go race detector and may panic. See issue #31.
//
//   - NewAPIClient falls back to http.DefaultClient (a process-wide
//     singleton with no Timeout) when cfg.HTTPClient is nil. A hung
//     server then hangs the caller indefinitely, and any mutation of
//     http.DefaultClient leaks into every other package in the process.
//     See issue #32.
//
// The helpers below provide safer alternatives without modifying the
// generated code. The underlying generated APIs (AddDefaultHeader,
// NewAPIClient) still exist and are still racy / shared-singleton; the
// recommended consumer pattern is documented in DEVELOPING.md.

// DefaultClientTimeout is applied to the *http.Client that
// NewAPIClientWithDefaults constructs when the caller does not supply
// their own. It is deliberately generous because the API includes
// long-running push/pull operations; callers with tighter budgets should
// construct an *http.Client directly and assign it to cfg.HTTPClient.
const DefaultClientTimeout = 5 * time.Minute

// NewAPIClientWithDefaults returns an APIClient configured with a
// non-shared *http.Client that carries a default timeout, instead of the
// process-wide http.DefaultClient singleton that the stock NewAPIClient
// installs. Pass a non-nil cfg.HTTPClient to skip the defaulting.
func NewAPIClientWithDefaults(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: DefaultClientTimeout}
	}
	return NewAPIClient(cfg)
}

// headerMu serializes access to Configuration.DefaultHeader across all
// Configuration instances in the process. A per-Configuration mutex would
// require modifying the generated struct, which we deliberately do not do
// (see the file-level comment). The number of Configuration instances per
// process is typically 1, so the global guard is acceptable in practice.
var headerMu sync.RWMutex

// SetDefaultHeader writes a default-header entry under a read/write lock.
// Use this in place of cfg.AddDefaultHeader when there is any chance of
// concurrent reads or writes (e.g. a background goroutine refreshing a
// bearer token while requests are in flight).
//
// Note: the generated prepareRequest path still iterates DefaultHeader
// without acquiring this lock. The lock makes header writes safe relative
// to each other and relative to DefaultHeaders snapshots, but does NOT
// fully protect against a concurrent prepareRequest iteration. The robust
// consumer pattern is to call SetDefaultHeader only from the init phase,
// before NewAPIClient or NewAPIClientWithDefaults is invoked. See issue
// #31 for the upstream fix tracking.
func SetDefaultHeader(cfg *Configuration, key, value string) {
	headerMu.Lock()
	defer headerMu.Unlock()
	if cfg.DefaultHeader == nil {
		cfg.DefaultHeader = make(map[string]string)
	}
	cfg.DefaultHeader[key] = value
}

// DefaultHeaders returns a copy of cfg.DefaultHeader taken under a read
// lock. Pair it with SetDefaultHeader on writes to avoid the "concurrent
// map iteration and map write" runtime panic on the read side.
func DefaultHeaders(cfg *Configuration) map[string]string {
	headerMu.RLock()
	defer headerMu.RUnlock()
	out := make(map[string]string, len(cfg.DefaultHeader))
	for k, v := range cfg.DefaultHeader {
		out[k] = v
	}
	return out
}
