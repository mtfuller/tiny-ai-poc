---
applyTo: "**/*.go"
---

# Coding Standards — Starterpack Go CLI (Concise)

## Style & Naming
- Packages: lowercase, single word (logger, color, spinner, version).  
- Files: lowercase with underscores (user_service.go) or without (greet.go, calc.go).  
- Functions: Exported CamelCase, unexported camelCase.  
- Vars: descriptive; common short names ok (cmd, err, ctx, args).  
- Constants: CamelCase or UPPER_SNAKE_CASE.

## File Layout
Order: package → imports (stdlib, external, internal) → constants → types → init() → constructors → methods → helpers.

Example:
```go
package cmd

import (
    "fmt"
    "os"
    
    "github.com/spf13/cobra"
    
    "github.com/mtfuller/starterpack-go-cli/internal/color"
    "github.com/mtfuller/starterpack-go-cli/internal/logger"
)
```

## Errors
- Always check errors and add context: fmt.Errorf("...: %w", err).  
- Log errors with logger at appropriate level (logger.Error, logger.Warn).  
- Exit with os.Exit(1) for fatal errors in command execution.
- Use cobra.Command.RunE for commands that can fail; return errors rather than os.Exit.

## Logging
- Use custom logger from internal/logger with structured formatting.  
- Levels: DEBUG, INFO, WARN, ERROR (controlled via --verbose or --log-level flags).  
- Log important operations with context (logger.Info, logger.Debug).
- Use color package for user-facing output (color.Success, color.Error, color.Info, color.Warn).

## Cobra Commands (pattern)
1. Define cobra.Command with Use, Short, Long, Example fields.
2. Add flags in init() function with cmd.Flags() or cmd.PersistentFlags().
3. Implement Run or RunE function to execute command logic.
4. Register command with parent using AddCommand() in init().

Example:
```go
var greetCmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone",
    Long:  "A longer description of the greet command",
    Args:  cobra.MaximumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := "World"
        if len(args) > 0 {
            name = args[0]
        }
        color.Success("Hello, %s!", name)
        return nil
    },
}

func init() {
    rootCmd.AddCommand(greetCmd)
    greetCmd.Flags().BoolP("uppercase", "u", false, "Print in uppercase")
}
```

## Project Structure
- cmd/: Cobra commands (one file per command: greet.go, calc.go, etc.)
- internal/: CLI-specific logic (logger, color, spinner, version)
- pkg/: Reusable libraries that could be used outside this CLI
- main.go: Entry point that calls cmd.Execute()
- Avoid globals except for command definitions and package-level logger instances

## Flags & Arguments
- Use cobra.Command.Flags() for command-specific flags
- Use rootCmd.PersistentFlags() for global flags (--verbose, --log-level)
- Validate required flags with cmd.MarkFlagRequired()
- Use cobra argument validators (cobra.ExactArgs, cobra.MinimumNArgs, etc.)
- Bind flags to variables with cmd.Flags().StringVarP() or similar

## Review Checklist
- [ ] Tests pass (unit and integration)
- [ ] No hardcoded values (use flags/args)
- [ ] Errors logged with context
- [ ] Proper exit codes (0 for success, 1 for errors)
- [ ] Input validated (args, flags)
- [ ] gofmt applied
- [ ] No sensitive data in logs
- [ ] Exported functions commented
- [ ] README updated if new command/flag added
- [ ] Help text (Short, Long) added to commands

## Patterns
- Keep cmd/ files focused on CLI interface; move business logic to internal/ or pkg/
- Use dependency injection where possible (pass logger, config to functions)
- Define interfaces in consumer packages when needed
- Use context.Context for cancellation in long-running operations
- Colored output for user messages; structured logging for debugging
- Test commands using integration tests that execute the full CLI

This preserves the project's conventions (Cobra, custom logger, color output) and focuses on clear, testable, and user-friendly CLI code.
