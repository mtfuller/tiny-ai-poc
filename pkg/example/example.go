package example

import (
	"fmt"
	"time"

	"github.com/mtfuller/starterpack-go-cli/internal/color"
	"github.com/mtfuller/starterpack-go-cli/internal/logger"
	"github.com/mtfuller/starterpack-go-cli/internal/spinner"
)

// ProcessData is an example function that demonstrates spinner and logging
func ProcessData(items []string, verbose bool) error {
	if verbose {
		logger.SetLevel(logger.DEBUG)
	}

	logger.Info("Starting data processing")
	logger.Debug("Processing %d items", len(items))

	s := spinner.New("Processing items...")
	s.Start()

	// Simulate processing
	time.Sleep(2 * time.Second)

	s.Stop()

	color.Success("Successfully processed %d items", len(items))
	return nil
}

// Greet is a simple greeting function
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Calculate performs a simple calculation
func Calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}
