package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rs-jensen/browser-cleaner/cleaner"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove browser data",
	RunE: func(cmd *cobra.Command, args []string) error {
		targets := cleaner.ResolveBrowsers(browser)
		if len(targets) == 0 {
			color.Yellow("No supported browsers found.")
			return nil
		}

		summary := cleaner.NewSummary()

		for _, b := range targets {
			result, err := cleaner.Clean(b, dryRun, verbose)
			if err != nil {
				color.Red("[%s] Error: %v", b.Name, err)
				continue
			}
			summary.Add(result)
		}

		fmt.Println()
		summary.Print()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
