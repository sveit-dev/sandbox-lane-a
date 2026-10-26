# sandbox-lane-a

A test fixture for the Sveit integration and end-to-end tests. It is not a
product and nothing depends on it.

- `slug`, `cart`, `pager`: small packages with table-driven tests.
- `protected/`: a config file and the `MaxItems` constant that the rest of the
  code reads. The end-to-end tests treat this directory as protected.
- `testdata/KNOWN_BUG.md`: a deliberate bug in an untested function, used by
  the qa end-to-end case.

`make build`, `make test` and `make lint` run `go build`, `go test` and
`go vet` over the module. CI runs build and test on every push and pull request.
