package cmd

import (
	"github.com/mtfuller/starterpack-go-cli/internal/color"
	"github.com/mtfuller/starterpack-go-cli/pkg/example"
	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process data with spinner animation",
	Long:  `Demonstrates spinner animation and logging during a long-running operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		items := []string{"item1", "item2", "item3", "item4", "item5"}
		
		err := example.ProcessData(items, verbose)
		if err != nil {
			color.Error("Failed to process data: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(processCmd)
}
