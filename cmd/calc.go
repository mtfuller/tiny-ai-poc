package cmd

import (
	"fmt"

	"github.com/mtfuller/starterpack-go-cli/internal/color"
	"github.com/mtfuller/starterpack-go-cli/internal/logger"
	"github.com/mtfuller/starterpack-go-cli/pkg/example"
	"github.com/spf13/cobra"
)

var operation string

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc [num1] [num2]",
	Short: "Perform calculations",
	Long:  `Demonstrates argument parsing with flags. Supports add, subtract, multiply, and divide operations.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		fmt.Sscanf(args[0], "%d", &a)
		fmt.Sscanf(args[1], "%d", &b)

		logger.Debug("Calculating: %d %s %d", a, operation, b)

		result, err := example.Calculate(a, b, operation)
		if err != nil {
			color.Error("Calculation error: %v", err)
			return
		}

		color.Success("Result: %d", result)
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
	calcCmd.Flags().StringVarP(&operation, "operation", "o", "add", "operation to perform (add, subtract, multiply, divide)")
}
