# DSL Transformer

[![CI](https://github.com/tommed/dsl-transformer/actions/workflows/ci.yml/badge.svg)](https://github.com/tommed/dsl-transformer/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/tommed/dsl-transformer/branch/main/graph/badge.svg)](https://codecov.io/gh/tommed/dsl-transformer)

---

## About

`dsl-transformer` is a lightweight, embeddable data transformation engine designed for structured data (JSON, maps, structs). Transformations are defined using a DSL (JSON or HCL-based), making it suitable for use cases like:

- Event stream manipulation
- API response mutation
- ETL pipelines
- Testing tools and CLI automation

It is:
- ⚡ Minimal
- 🟣 Composable
- 🟩 Fully testable

## DSL Specification

See [SPEC.md](./SPEC.md) for a specification of the DSL.

---

## Features (WIP)
- [x] Declarative `set` operations
- [ ] Support for `map`, `delete`, `merge`
- [ ] Conditionals
- [ ] CLI for local testing
- [ ] Embeddable Go SDK
- [ ] HCL-powered syntax option
- [ ] Serverless runtime compatibility
- [ ] OpenTelemetry instrumentation

---

## Example

### `examples/simplest.json`

```json
{
  "version": 1,
  "instructions": [
    {"op": "set", "key": "greeting", "value": "hello world"}
  ]
}
```

## Run

```bash
echo '{"foo":"bar"}' | go run ./cmd/transformer-cli examples/simplest.json
```

Output:

```json
{
  "foo": "bar",
  "greeting": "hello world"
}

```

## Development

### Testing

```bash
make test         # Short tests
make test-full    # Full tests
make coverage     # Coverage report (HTML)
make lint-install # Install lint prerequisites
make lint         # Run static analysis
```

### CLI

```bash
make cli-macos
```

## Status

See [status.md](./status.md) for up-to-date CI, coverage, and project health.

## License

[MIT](./LICENSE)