package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rs-jensen/browser-cleaner/cleaner"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan browser data without removing",
	RunE: func(cmd *cobra.Command, args []string) error {
		targets := cleaner.ResolveBrowsers(browser)
		if len(targets) == 0 {
			color.Yellow("No supported browsers found.")
			return nil
		}

		for _, b := range targets {
			report, err := cleaner.Scan(b, verbose)
			if err != nil {
				color.Red("[%s] Error: %v", b.Name, err)
				continue
			}
			report.Print()
			fmt.Println()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
