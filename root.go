package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	dryRun  bool
	verbose bool
	browser string
)

var rootCmd = &cobra.Command{
	Use:   "browser-cleaner",
	Short: "Privacy tool for cleaning browser data",
	Long:  "Scans and removes cookies, cache, and session data from installed browsers.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "preview without deleting")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "detailed output")
	rootCmd.PersistentFlags().StringVarP(&browser, "browser", "b", "all", "target browser (chrome, firefox, edge, all)")

	fmt.Println()
	color.Cyan("🧹 Browser Cleaner — Privacy Tool")
	fmt.Println()
}
