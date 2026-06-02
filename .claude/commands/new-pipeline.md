# /new-pipeline — Design and scaffold a new agent Pipeline

Create a multi-step agent pipeline for a specific task.

## Arguments
Pass a short description of the pipeline's purpose, e.g. `/new-pipeline summarize a document`

## What to do

1. **Plan the pipeline steps** — think through what reasoning stages are needed for small models to reach the right output. Each stage should be a single, focused operation (classify, extract, summarize, validate, etc.). Describe the plan to the user and ask for confirmation before writing code.

2. **Create `pkg/agent/<task>_pipeline.go`**:
   - A constructor `New<Task>Pipeline(client, model string) *agent.Pipeline` 
   - Each step as a separate `Step` implementation in the same file (or split to `<step>_step.go` if complex)
   - A private `build<Step>Prompt` function per step

3. **Create `pkg/agent/<task>_pipeline_test.go`**:
   - Mock Ollama client
   - Table-driven tests for the full pipeline with representative inputs
   - Tests for each individual step

4. **Wire it into a command** in `cmd/` if one doesn't already exist.

## Small-model pipeline design rules

- **One job per step** — a step either transforms, classifies, extracts, or validates. Never all at once.
- **Constrain every output** — tell the model exactly what format to return. Parse and validate it. Retry once on parse failure.
- **Summarize, don't truncate** — if prior context is too long for the next step's prompt, summarize it in a prior step.
- **Explicit chain-of-thought** — if the task needs reasoning, add a dedicated "think" step that writes to `state.Metadata["reasoning"]`, then a separate "answer" step that reads from it.
- **Fail fast** — if a step gets back malformed output twice, return an error rather than guessing.
