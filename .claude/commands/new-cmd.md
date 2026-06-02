# /new-cmd — Scaffold a new Cobra CLI command

Add a new subcommand to `tiny-ai`.

## Arguments
Pass the command name as the argument, e.g. `/new-cmd summarize`

## What to do

1. Create `cmd/<name>.go`:
   - Define `var <name>Cmd = &cobra.Command{Use, Short, Long, Example, RunE}`
   - Declare and bind flags in `init()` using `<name>Cmd.Flags()`
   - Register with `rootCmd.AddCommand(<name>Cmd)` in `init()`
   - Business logic goes in `pkg/` or `internal/` — keep `cmd/` thin

2. Add integration test to `tests/integration_test.go`:
   - Test the happy path
   - Test `--help` works
   - Test one error case (bad args or missing flag)

3. Update `README.md` commands table with the new command.

## Command template

```go
package cmd

import (
    "github.com/mtfuller/tiny-ai-poc/internal/color"
    "github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
    Use:     "mycommand [args]",
    Short:   "One-line description",
    Long:    `Longer description shown in --help.`,
    Example: `  tiny-ai mycommand --flag value`,
    RunE: func(cmd *cobra.Command, args []string) error {
        color.Info("Running mycommand...")
        return nil
    },
}

func init() {
    rootCmd.AddCommand(myCmd)
    myCmd.Flags().StringP("flag", "f", "default", "flag description")
}
```

## Conventions
- Use `RunE` (not `Run`) so errors propagate cleanly
- Use `color.Info/Success/Error/Warn` for user-facing output
- Use `logger.Debug/Info` for diagnostic output
- Validate args with `cobra.ExactArgs`, `cobra.MinimumNArgs`, etc.
- Flags with `--ollama-url` and `--model` should always default to `http://localhost:11434` and `llama3.2:latest`
