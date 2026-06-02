package cmd

import (
	"os"

	"github.com/mtfuller/starterpack-go-cli/internal/color"
	"github.com/mtfuller/starterpack-go-cli/internal/logger"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	logLevel string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "starterpack-go-cli",
	Short: "A state-of-the-art Go CLI application template",
	Long: color.Bold("starterpack-go-cli") + ` is a comprehensive Go CLI application template
that includes many features out-of-the-box:
  • Argument parsing with Cobra
  • Structured logging
  • Colored text output
  • Spinner animations
  • Version command
  • Help command
  • Unit and integration tests

This template helps developers quickly bootstrap a professional Go CLI application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Configure logger based on flags
		if verbose {
			logger.SetLevel(logger.DEBUG)
		} else {
			logger.SetLevel(logger.ParseLogLevel(logLevel))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		color.Error("Error: %v", err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output (debug level)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "set log level (debug, info, warn, error)")
}
