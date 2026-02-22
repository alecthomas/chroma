
# Contributing to Chroma

Thank you for wanting to contribute.

## Quick Start

- **Clone:** Use the normal fork-and-branch workflow (fork, create a topic branch).
- **Build & test:** Run `go test ./...` to run tests across the repository.
- **Just:** This repository uses `just` for helper commands; run `just -l` to list available targets and see `./Justfile` for details.

## Testing and Running the Playground

- **Run the web playground (local):** from the repo root run:

```sh
go run -C ./cmd/chromad . --csrf-key=securekey
```

- **Lexer tests:** To run only lexers tests:

```sh
go test ./lexers
```

- **Regenerating lexer expected outputs:** When you update lexers and add new `*.actual` testdata files, regenerate expected outputs by running:

```sh
RECORD=true go test ./lexers
```

This mirrors the behaviour described in [lexers/README.md](lexers/README.md).

## Developing Lexers

- Prefer XML lexer definitions in `lexers/embedded/*.xml` unless a lexer needs custom code. Use the `_tools/pygments2chroma_xml.py` helper to convert from Pygments when possible; see [lexers/README.md](lexers/README.md) for details and test patterns.
