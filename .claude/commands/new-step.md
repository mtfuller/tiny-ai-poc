# /new-step — Scaffold a new agent pipeline Step

Create a new `Step` implementation in `pkg/agent/` for this project.

## Arguments
Pass the step name as the argument, e.g. `/new-step summarize`

## What to do

1. Create `pkg/agent/<name>_step.go` with:
   - A struct holding an Ollama client interface and model name
   - `Name() string` returning the step name
   - `Run(ctx context.Context, state *State) error` implementing the logic
   - A `ChatClient` or `GenerateClient` interface so tests can mock without Ollama

2. Create `pkg/agent/<name>_step_test.go` with:
   - A `Mock<Client>` struct implementing the interface
   - Table-driven tests covering the happy path and at least one error case
   - Assertions on `state.Output` and any `state.Metadata` the step sets

3. If the step requires prompt construction logic, put it in a private `build<Name>Prompt(state *State) string` function and test it separately.

4. Update `CLAUDE.md` "Architecture" section to mention the new step.

## Conventions
- System prompt: one sentence, role-defining, no padding
- User prompt: minimal — only what the model strictly needs
- Always constrain output format in the prompt (e.g., "respond with JSON only")
- Never call `time.Sleep` or add artificial delays
- Step must be stateless except through `state`
