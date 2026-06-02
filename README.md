# starterpack-go-cli

A state-of-the-art Go CLI application template that includes many features out-of-the-box to help developers quickly bootstrap a professional command-line application.

## Features

✨ **Out-of-the-box features:**

- 🎯 **Argument Parsing**: Built with [Cobra](https://github.com/spf13/cobra) for robust command-line interface
- 📝 **Structured Logging**: Custom logger with multiple log levels (DEBUG, INFO, WARN, ERROR)
- 🎨 **Colored Text Output**: ANSI color support for beautiful terminal output
- ⏳ **Spinner Animations**: Visual feedback for long-running operations
- 📦 **Version Command**: Built-in version management with build metadata
- ❓ **Help Command**: Auto-generated help documentation for all commands
- ✅ **Unit Tests**: Comprehensive unit tests for all packages
- 🧪 **Integration Tests**: End-to-end integration tests for CLI commands
- 🔨 **Taskfile**: Easy build, test, and run commands with [Task](https://taskfile.dev)

## Quick Start

### Prerequisites

- Go 1.21 or higher
- [Task](https://taskfile.dev) (optional, for build automation)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/mtfuller/starterpack-go-cli.git
cd starterpack-go-cli
```

2. Build the application:
```bash
task build
```

3. Run the application:
```bash
./starterpack-go-cli --help
```

## Usage

### Available Commands

#### Version Command
Display version information:
```bash
./starterpack-go-cli version
```

Short version output:
```bash
./starterpack-go-cli version --short
```

#### Greet Command
Simple greeting with colored output:
```bash
./starterpack-go-cli greet Alice
```

#### Calc Command
Perform calculations with different operations:
```bash
./starterpack-go-cli calc 10 5 --operation add
./starterpack-go-cli calc 10 5 --operation subtract
./starterpack-go-cli calc 10 5 --operation multiply
./starterpack-go-cli calc 10 5 --operation divide
```

#### Process Command
Demonstrates spinner animation and logging:
```bash
./starterpack-go-cli process
```

### Global Flags

- `-v, --verbose`: Enable verbose output (debug level logging)
- `-l, --log-level`: Set log level (debug, info, warn, error)
- `-h, --help`: Display help information

### Examples with Flags

Enable verbose logging:
```bash
./starterpack-go-cli greet World --verbose
```

Set specific log level:
```bash
./starterpack-go-cli process --log-level debug
```

## Development

### Running Tests

Run all tests:
```bash
task test
```

Run only unit tests:
```bash
task test-unit
```

Run only integration tests:
```bash
task test-integration
```

Generate coverage report:
```bash
task coverage
```

### Building

Build the binary:
```bash
task build
```

Install to GOPATH/bin:
```bash
task install
```

### Project Structure

```
.
├── cmd/                    # Command definitions
│   ├── root.go            # Root command
│   ├── version.go         # Version command
│   ├── greet.go           # Example greet command
│   ├── calc.go            # Example calc command
│   └── process.go         # Example process command
├── internal/              # Internal packages
│   ├── color/             # Colored text utilities
│   ├── logger/            # Structured logging
│   ├── spinner/           # Spinner animations
│   └── version/           # Version management
├── pkg/                   # Public packages
│   └── example/           # Example business logic
├── tests/                 # Integration tests
├── main.go               # Application entry point
├── Taskfile.yml          # Build and test automation
└── README.md             # This file
```

## Adding New Commands

To add a new command, create a new file in the `cmd/` directory:

```go
package cmd

import (
    "github.com/spf13/cobra"
    "github.com/mtfuller/starterpack-go-cli/internal/color"
)

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description of my command",
    Run: func(cmd *cobra.Command, args []string) {
        color.Success("My command executed!")
    },
}

func init() {
    rootCmd.AddCommand(myCmd)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
