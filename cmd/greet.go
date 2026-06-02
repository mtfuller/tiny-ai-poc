package cmd

import (
	"github.com/mtfuller/starterpack-go-cli/internal/color"
	"github.com/mtfuller/starterpack-go-cli/internal/logger"
	"github.com/mtfuller/starterpack-go-cli/pkg/example"
	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greet someone by name",
	Long:  `A simple command that demonstrates argument parsing and colored output.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		logger.Debug("Greeting user: %s", name)
		
		greeting := example.Greet(name)
		color.Success(greeting)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
