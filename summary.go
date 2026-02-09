package cleaner

import (
	"fmt"

	"github.com/fatih/color"
)

type Summary struct {
	Results []*CleanResult
}

func NewSummary() *Summary {
	return &Summary{}
}

func (s *Summary) Add(r *CleanResult) {
	s.Results = append(s.Results, r)
}

func (s *Summary) Print() {
	color.Cyan("━━━ Summary ━━━━━━━━━━━━━━━━━━━━━━━━━")

	var totalFiles int
	var totalBytes int64
	var totalErrors int

	for _, r := range s.Results {
		fmt.Printf("  %-10s %4d files  %8s  ", r.Browser, r.FilesRemoved, formatBytes(r.BytesFreed))
		if r.Errors > 0 {
			color.Red("%d errors", r.Errors)
		} else {
			color.Green("ok")
		}
		totalFiles += r.FilesRemoved
		totalBytes += r.BytesFreed
		totalErrors += r.Errors
	}

	fmt.Println()
	color.White("  Total: %d files, %s freed", totalFiles, formatBytes(totalBytes))
	if totalErrors > 0 {
		color.Red("  Errors: %d", totalErrors)
	}
	fmt.Println()
}
