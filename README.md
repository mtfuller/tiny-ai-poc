# tiny-ai

A proof-of-concept general-purpose AI assistant powered by small, locally-run language models via [Ollama](https://ollama.com).

## The Thesis

Large cloud models are powerful but expensive. Tiny models (Llama 3.2 1B/2B, etc.) are fast, free, and fully private — but they require more engineering to produce reliable results. This project explores whether a combination of:

- **Structured agent pipelines** (explicit, inspectable reasoning steps)
- **Programmatic context management** (selective memory, token budgets)
- **Constrained prompt engineering** (tight output schemas, chain-of-thought hints)

…can close enough of the capability gap to build a genuinely useful tool.

## Prerequisites

- Go 1.21+
- [Ollama](https://ollama.com) running locally (`ollama serve`)
- A small model pulled: `ollama pull llama3.2:latest`
- [Task](https://taskfile.dev) (optional)

## Quick Start

```bash
# Build
task build

# Start a chat session
./tiny-ai chat

# Use a specific model
./tiny-ai chat --model llama3.2:1b

# Point at a remote Ollama instance
./tiny-ai chat --model llama3.2:latest --ollama-url http://my-server:11434
```

## Commands

| Command | Description |
|---------|-------------|
| `chat` | Interactive multi-turn chat session |
| `version` | Print version / build info |

Global flags: `--verbose / -v`, `--log-level / -l`

## Project Structure

```
.
├── cmd/                    # Cobra CLI commands
│   ├── root.go
│   ├── chat.go             # Interactive chat command
│   └── version.go
├── pkg/
│   ├── ollama/             # Ollama HTTP client (Chat, Generate, ListModels)
│   └── agent/              # Pipeline primitives: Step, State, Pipeline
├── internal/
│   ├── color/              # ANSI terminal colors
│   ├── logger/             # Structured leveled logger
│   ├── spinner/            # Progress spinner
│   └── version/            # Build-time version info
├── tests/                  # Integration tests
├── Taskfile.yml
└── CLAUDE.md               # AI dev context
```

## Development

```bash
task test            # all tests
task test-unit       # unit tests only
task test-integration
task coverage        # HTML coverage report
task lint            # go vet + gofmt check
```

## Design Principles

1. **Explicit over implicit** — every reasoning step is a named `Step` in a `Pipeline`, not a magic prompt.
2. **Token budget discipline** — context is constructed programmatically; nothing is appended naively.
3. **Small model assumptions** — prompts are short, output formats are constrained, multi-step > single-shot.
4. **Local first** — no cloud APIs required; works entirely with Ollama on localhost.

## License

MIT
