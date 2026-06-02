---
applyTo: "**/*_test.go"
---
## Philosophy
- Test observable behavior, not implementation.
- Arrange → Act → Assert; keep tests isolated and fast.
- Prioritize critical paths; aim overall >80% coverage.

## Layout
tests/
- integration_test.go — full CLI command execution via os/exec  
- Unit tests — *_test.go in same package as source (internal/, pkg/)

## Code style 
- Unit tests: *_test.go in same package for white-box access when needed.  
- Integration tests: tests/integration_test.go for end-to-end CLI commands.  
- Use table-driven tests and testify for assertions.  
- Aim >80% coverage; mock external deps (filesystem, network).

## Unit testing (guidelines)
- Keep tests in the same package (e.g., internal/logger/logger_test.go).
- Use table-driven tests for variants and edge cases.
- Mock external deps (filesystem operations, external APIs) for isolation.
- Use bytes.Buffer to capture and test logger output.
- Test both success and error paths; assert error messages.
- Test flag parsing and validation logic in commands.

## Integration testing (guidelines)
- Use os/exec.Command("go", "run", "../main.go", ...) to test full CLI execution.
- Capture stdout/stderr with bytes.Buffer attached to cmd.Stdout and cmd.Stderr.
- Test command success (exit code 0) and failure cases (exit code 1).
- Test all major commands and flag combinations (e.g., --short, --verbose).
- Verify output contains expected strings, colored output, and error messages.
- Use t.Cleanup() for temporary files or state cleanup.

## Best practices
- Table-driven tests for coverage and clarity.
- Cover edge cases (invalid flags, missing arguments, empty input).
- Create small test helpers (captureOutput, runCommand).
- Clean up resources (temp files, test artifacts) with t.Cleanup().
- Test terminal output formatting (colors, spinners) by checking ANSI codes.
- Test logger levels (DEBUG, INFO, WARN, ERROR) with buffer capture.

## Running tests
- task test        # all tests
- task test-unit   # unit tests only
- task test-integration # integration tests only
- task coverage    # opens coverage report
- go test ./... -v
- go test -race ./...

## Coverage targets
- Commands >85%, Business logic >85%, Utilities >80%, Overall >80%.

## CI
- Run on PRs and main commits: tests, coverage gate, linters, race detector.

Keep tests small, deterministic, and focused on CLI behavior.
