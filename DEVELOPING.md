# Project Development

For general information about contributing changes, see the
[Contributor Guidelines](https://github.com/datadatdat/.github/blob/master/CONTRIBUTING.md).

## How it Works

The project is entirely generated via the [OpenAPI Generator](https://openapi-generator.tech/). The resulting
files are committed to the repository so they can be easily imported via go modules. To updated the generated
code, you will need to:

1. Update `datadatdat.yml` with the specification from the `datadatdat-server` repository.
2. Run `generate`. This requires docker, and likely work on windows systems. If necessary, you can run the openapi
   generator by hand, but you will be responsible for ensuring the right options are specified.
3. If the new API spec contains renamed or deleted objects, you will need to manually delete those objects.
4. Run `go build` to properly set the go version in the `go.mod` file.

If you update to a later version of the openapi generator, you will need to update the mustache templates in
`templates`, which include a change to `api.mustache` to pass header parameters as JSON instead of strings.

## Building

Run `go build` to build the package.

## Testing

Because the code is completely auto-generated, there are no dedicated tests for the client, as we'd just be testing
the operation of the openapi generator.

## Releasing

Since it is delivered as a go module, releasing requires only creating a new tag with the appropriate version
string.
