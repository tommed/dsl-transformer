<!--suppress HtmlDeprecatedAttribute -->
<p align="right">
    <a href="https://github.com/tommed" title="See Project Ducto">
        <img src="./assets/ducto-logo-small.png" alt="A part of Project Ducto"/>
    </a>
</p>

# DSL Transformer

[![CI](https://github.com/tommed/dsl-transformer/actions/workflows/ci.yml/badge.svg)](https://github.com/tommed/dsl-transformer/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/tommed/dsl-transformer/branch/main/graph/badge.svg)](https://codecov.io/gh/tommed/dsl-transformer)

---

## About

<p align="center">
  <img alt="Graphical representation of Ducto manipulate streaming data in a system of pipes" 
       src="./assets/ducto-representation-small.png"/>
</p>

`dsl-transformer` is a lightweight, embeddable data transformation engine designed for structured data (JSON, maps, structs). Transformations are defined using a DSL (JSON or HCL-based), making it suitable for use cases like:

- Event stream manipulation
- API response mutation
- ETL pipelines
- Testing tools and CLI automation

It is:
- ⚡ Minimal
- 🟣 Composable
- 🟩 Fully testable

`dsl-transformer` is a part of the larger [Ducto project](https://github.com/tommed), combining many interesting practices together including Feature Flagging.

![Topology Diagram](./assets/topology-medium.png)

## DSL Specification

See [doc/SPEC.md](docs/specs.md) for a specification of the DSL.

---

## Features (WIP)
- [x] CLI for local testing
- [x] Declarative `set` and `copy` operations
- [x] Support for `map`, `delete`, `merge`
- [x] Linter included for instruction validation
- [ ] Conditionals
- [ ] Embeddable Go SDK
- [ ] Input can be JSON or YAML
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

## Lint

```bash
go run ./cmd/transformer-cli lint examples/simplest.json
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

See [status.md](docs/status.md) for up-to-date CI, coverage, and project health.

## License

- Code is all licensed under [MIT](./LICENSE)
- Logos/Illustrations are Copyright 2025 Tom Medhurst, all rights reserved.