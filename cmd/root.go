package cmd

import (
	"os"

	"github.com/mtfuller/tiny-ai-poc/internal/color"
	"github.com/mtfuller/tiny-ai-poc/internal/logger"
	"github.com/spf13/cobra"
)

var (
	verbose  bool
	logLevel string
)

var rootCmd = &cobra.Command{
	Use:   "tiny-ai",
	Short: "A general-purpose AI tool powered by small local LLMs",
	Long: color.Bold("tiny-ai") + ` is a proof-of-concept AI assistant powered by small,
locally-run language models via Ollama (e.g. Llama 3.2 2B).

Rather than relying on large cloud models, tiny-ai uses structured agent pipelines,
explicit context management, and programmatic prompt engineering to get useful results
from compact, fast, inexpensive models.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logger.SetLevel(logger.DEBUG)
		} else {
			logger.SetLevel(logger.ParseLogLevel(logLevel))
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		color.Error("Error: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output (debug level)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "set log level (debug, info, warn, error)")
}
