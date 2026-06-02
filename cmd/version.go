package cmd

import (
	"fmt"

	"github.com/mtfuller/tiny-ai-poc/internal/color"
	"github.com/mtfuller/tiny-ai-poc/internal/version"
	"github.com/spf13/cobra"
)

var short bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print the version, commit hash, and build date.`,
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			fmt.Println(version.GetShortVersion())
		} else {
			fmt.Println(color.Bold("tiny-ai"))
			fmt.Printf("Version:    %s\n", color.Cyan(version.Version))
			fmt.Printf("Commit:     %s\n", version.Commit)
			fmt.Printf("Built:      %s\n", version.BuildDate)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&short, "short", "s", false, "print only the version number")
}
