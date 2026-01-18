# advance

A minimal, local-first CLI tool for managing todos and advancing work from the terminal.

## Overview

Advance is a small, fast CLI for tracking todo items locally using a single JSON file (default: `~/.advance.json`). It aims to be simple, scriptable, and easy to extend.

## Installation

Install the CLI with Go (recommended):

```
go install github.com/tgenericx/advance/cmd@latest
```

Or build from source:

```
git clone https://github.com/tgenericx/advance.git
cd advance
go build -o advance ./cmd
```

## Usage

Default data file: `~/.advance.json` (override with `--datafile`).

Common commands:

```
# Add one or more todos
advance add "Write docs" "Fix bug"

# List all todos
advance list

# Mark a todo as done by ID
advance done 2
```

Notes:
- IDs are integers assigned to items. Use `advance list` to find IDs.
- `done` marks an item completed and persists the change.
- `advance add` accepts multiple arguments; items with spaces should be quoted.

## Data format

The tool stores todos as JSON in the data file. Items have an `ID`, `Text`, and `Done` boolean. The implementation uses a simple file-backed store located in `todo/filestore.go`.

## Development

- The CLI uses Cobra (`cmd/`) and a small `todo` package (`todo/`).
- Run tests with:

```
go test ./...
```

- Lint and vet as part of CI (recommended), for example:

```
go vet ./...
golangci-lint run
```

## Contributing

Contributions are welcome.

- Open issues for bugs or feature requests.
- Send PRs with tests and a clear description of changes.
- Follow Go formatting and linting standards.
