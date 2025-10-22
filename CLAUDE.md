# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based Buildkite plugin that demonstrates on-demand downloading and execution of pre-compiled binaries from GitHub releases. The plugin annotates Buildkite builds with custom messages.

## Development Commands

All Go development commands should be run from the `src/` directory:

- **Build and test**: `make test` (downloads dependencies, runs tests with race detection and coverage)
- **Run tests for CI**: `make test-ci` (generates coverage artifacts in `artifacts/` directory)
- **Download dependencies**: `make mod`
- **Ensure dependencies are tidy**: `make ensure-deps` (runs `go mod tidy` and checks for uncommitted changes)
- **Lint code**: Use `golangci-lint run -v --timeout=2m` (requires golangci-lint to be installed)

For shell script testing:

- **Run BATS tests**: `bats tests/download.bats` (requires BATS testing framework)

## Architecture

The plugin follows a clean architecture pattern:

### Core Components

- **`main.go`**: Entry point that orchestrates the plugin execution
- **`plugin/example.go`**: Main plugin logic with dependency injection via interfaces
- **`plugin/config.go`**: Configuration management using environment variables
- **`buildkite/agent.go`**: Buildkite agent interaction and command execution
- **`buildkite/log.go`**: Logging utilities

### Key Interfaces

- **`ConfigFetcher`**: Abstracts configuration retrieval (implemented by `EnvironmentConfigFetcher`)
- **`Agent`**: Abstracts Buildkite agent operations (implemented by `buildkite.Agent`)

### Plugin Flow

1. Configuration is fetched from environment variables with prefix `BUILDKITE_PLUGIN_EXAMPLE_GO`
2. The plugin annotates the Buildkite build with the provided message
3. Commands are executed via `buildkite-agent` CLI tool with proper signal handling

### Distribution Mechanism

The plugin uses a two-stage execution model:

- **Stage 1**: Bash script (`hooks/command`) downloads the appropriate binary for the current architecture
- **Stage 2**: Downloaded Go binary executes the actual plugin logic

## Testing Strategy

- **Go tests**: Standard Go testing with mocks generated using interfaces
- **Shell tests**: BATS framework for testing the download logic
- **CI**: GitHub Actions with Go checks, linting, and plugin-specific validation

## Configuration

The plugin expects a `message` parameter via environment variable `BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE`.00~## MCP Rules

- When using Context7 maintain a file named library.md to store a Library IDs that you search for and before searching make sure that you check the file and use the library ID already available. Otherwise search for it.01~

## MCP Rules

- When using Context7 maintain a file named library.md to store a Library IDs that you search for and before searching make sure that you check the file and use the library ID already available. Otherwise search for it.
