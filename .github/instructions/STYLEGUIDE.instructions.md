---
applyTo: "**/*.go"
---

# Coding Standards — tiny-ai-poc

## Style & Naming
- Packages: lowercase, single word (`ollama`, `agent`, `logger`).
- Files: lowercase with underscores or without (`chat.go`, `summarize_step.go`).
- Functions: exported CamelCase, unexported camelCase.
- Constants: CamelCase or UPPER_SNAKE_CASE.

## File Layout
Order: `package` → imports (stdlib, external, internal) → constants → types → constructors → methods → helpers.

```go
package agent

import (
    "context"
    "fmt"

    "github.com/mtfuller/tiny-ai-poc/pkg/ollama"
    "github.com/mtfuller/tiny-ai-poc/internal/logger"
)
```

## Errors
- Always wrap: `fmt.Errorf("context: %w", err)`.
- Use `logger.Error` / `logger.Warn` for operational errors.
- Commands use `RunE` and return errors; root handler calls `os.Exit(1)`.

## Ollama Client Usage
- Always pass `context.Context` as the first argument.
- Keep `Stream: false` until streaming is explicitly required.
- Construct an interface (`ChatClient`) in consuming packages so tests can mock without a live Ollama.

## Agent Pipeline Pattern
1. Define a struct that holds dependencies (Ollama client, model name, options).
2. Implement `Name() string` and `Run(ctx context.Context, state *agent.State) error`.
3. Read from `state.Input` or `state.Memory`, write results to `state.Output` or `state.Metadata`.
4. Register the step in a `Pipeline` from the calling `cmd/` layer.

```go
type MyStep struct {
    client ChatClient
    model  string
}

func (s *MyStep) Name() string { return "my-step" }

func (s *MyStep) Run(ctx context.Context, state *agent.State) error {
    resp, err := s.client.Generate(ctx, ollama.GenerateRequest{
        Model:  s.model,
        Prompt: buildPrompt(state),
        System: "...",
    })
    if err != nil {
        return fmt.Errorf("generate: %w", err)
    }
    state.Output = resp.Response
    return nil
}
```

## Prompt Engineering Rules
- System prompts: one clear role sentence, no fluff.
- User prompts: include only what the model needs — no extra context.
- Always constrain output format explicitly (e.g., "respond with only valid JSON", "one sentence only").
- For chain-of-thought, use a separate Step rather than asking the model to think in the same turn.

## Review Checklist
- [ ] Tests pass (`task test`)
- [ ] Ollama calls use `context.Context`
- [ ] No hardcoded model names (use flags/config)
- [ ] Step has a unit test with mocked client
- [ ] Prompt is short and output-constrained
- [ ] Errors wrapped with context
- [ ] `gofmt` applied (`task lint`)
- [ ] README / CLAUDE.md updated if new command or Step added
