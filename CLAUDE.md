# CLAUDE.md — tiny-ai-poc

## Project Purpose

This is a Go CLI tool (`tiny-ai`) that demonstrates how to build a useful AI assistant using small, locally-run models (Llama 3.2 1B/2B, etc.) via Ollama. The core hypothesis is that structured agent pipelines + programmatic context management can compensate for the capability limits of small models.

## Architecture

```
cmd/          CLI entry points (Cobra commands)
pkg/ollama/   Ollama HTTP client — Chat, Generate, ListModels
pkg/agent/    Agent pipeline primitives — Step, State, Pipeline
internal/     CLI utilities (color, logger, spinner, version)
tests/        Integration tests
```

### Key Abstractions

**`pkg/agent`** — the heart of the project.
- `State` holds the mutable data flowing through a pipeline (input, output, memory, metadata).
- `Step` is an interface with a single `Run(ctx, *State) error` method.
- `Pipeline` composes steps in sequence and runs them against shared state.
- `MemoryEntry` with `Pinned bool` supports selective context inclusion.

**`pkg/ollama`** — thin HTTP client over the Ollama REST API.
- `Client.Chat()` — multi-turn conversation
- `Client.Generate()` — single-turn completion
- `Client.ListModels()` — available local models

**`cmd/chat.go`** — wires Ollama client into an interactive REPL.

## Design Principles (enforce these in code reviews)

1. **Small-model-first prompting** — prompts must be short and explicit. Never rely on the model to infer structure. Always constrain output format.
2. **Explicit pipeline steps** — reasoning stages must be separate named `Step` implementations, not smashed into one big prompt.
3. **Token budget awareness** — before adding context to a prompt, check if it fits within the model's context window. Prefer summarization over truncation.
4. **No cloud dependencies** — everything runs via local Ollama. Do not add Anthropic/OpenAI SDKs unless the user explicitly asks.
5. **Testable steps** — every `Step` implementation must have a unit test. Mock the Ollama client with an interface.

## Development Commands

```bash
task build           # compile ./tiny-ai binary
task run             # go run main.go
task chat            # go run main.go chat
task test            # all tests
task test-unit       # internal/... pkg/...
task test-integration
task coverage
task lint            # go vet + gofmt
```

## Module Path

`github.com/mtfuller/tiny-ai-poc`

All internal imports use this prefix.

## Code Conventions

- New commands: `cmd/<name>.go`, register in `init()` with `rootCmd.AddCommand()`
- New agent steps: `pkg/agent/<name>_step.go`, implement `Step` interface
- New Ollama features: add to `pkg/ollama/client.go`
- CLI-specific helpers: `internal/<package>/`
- Tests live next to source (`*_test.go`) except integration tests in `tests/`

## Adding a New Agent Step

```go
// pkg/agent/summarize_step.go
package agent

import (
    "context"
    "fmt"
    "github.com/mtfuller/tiny-ai-poc/pkg/ollama"
)

type SummarizeStep struct {
    client *ollama.Client
    model  string
}

func (s *SummarizeStep) Name() string { return "summarize" }

func (s *SummarizeStep) Run(ctx context.Context, state *State) error {
    resp, err := s.client.Generate(ctx, ollama.GenerateRequest{
        Model:  s.model,
        Prompt: fmt.Sprintf("Summarize in one sentence:\n%s", state.Input),
        System: "You are a concise summarizer. Output only the summary sentence.",
    })
    if err != nil {
        return err
    }
    state.Output = resp.Response
    return nil
}
```

## Context Window Strategy

Small models (1B–3B params) typically have 4K–8K token context windows. The agent system should:
- Track token estimates per `MemoryEntry`
- Include `Pinned` entries unconditionally
- Fill remaining budget with most-recent non-pinned entries
- Summarize older turns rather than dropping them

This logic should live in a `ContextBuilder` in `pkg/agent/`.

## Testing the Ollama Client

Use an interface to mock the Ollama client in unit tests:

```go
type ChatClient interface {
    Chat(ctx context.Context, req ollama.ChatRequest) (*ollama.ChatResponse, error)
}
```

Implement `MockChatClient` in test files to avoid requiring a live Ollama instance during `task test-unit`.
