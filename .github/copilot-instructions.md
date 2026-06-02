You are a knowledgeable software engineer familiar with Go, small-model AI systems, and the Ollama API. Use the information below when assisting with this project.

## Project summary
- `tiny-ai` is a Go CLI tool that uses small local LLMs (via Ollama) with structured agent pipelines and programmatic context management to produce reliable results from compact models.
- Module: `github.com/mtfuller/tiny-ai-poc`
- Binary name: `tiny-ai`

## Layout
- `main.go` — calls `cmd.Execute()`
- `cmd/` — Cobra commands: `root.go`, `version.go`, `chat.go`
- `pkg/ollama/` — HTTP client for Ollama REST API (Chat, Generate, ListModels)
- `pkg/agent/` — Pipeline primitives: `Step`, `State`, `Pipeline`, `MemoryEntry`
- `internal/` — CLI utilities: `logger`, `color`, `spinner`, `version`
- `tests/` — integration tests

## Tech stack
- Cobra (CLI), custom leveled logger, testify, Task runner, Ollama REST API

## Core design principles
1. **Small-model-first** — prompts must be short, explicit, and output-constrained. Never rely on inference.
2. **Explicit pipelines** — reasoning = named `Step` chain, not a single prompt.
3. **Token budget** — context is constructed programmatically with a size budget.
4. **Local only** — Ollama on localhost; no cloud API keys required.
5. **Testable** — every Step is unit-testable via a mocked Ollama client interface.

## Key patterns
- New command: `cmd/<name>.go` with `cobra.Command`, register in `init()`.
- New pipeline step: `pkg/agent/<name>_step.go` implementing `Step` interface.
- Ollama calls: always pass `context.Context`; always set `Stream: false` for now.
- Errors: wrap with `fmt.Errorf("...: %w", err)`; use `logger` for debug output.

## Quality
- Unit tests for all `pkg/` and `internal/` code
- Integration tests in `tests/` for CLI commands
- Table-driven tests, testify assertions, >80% coverage goal
- Mock Ollama client with a `ChatClient` interface in tests

## Common tasks
- Build: `task build`
- Chat: `task chat`
- Test: `task test` / `task test-unit` / `task test-integration`
- Lint: `task lint`
