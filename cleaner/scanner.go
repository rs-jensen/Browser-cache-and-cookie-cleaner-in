package cleaner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type ScanReport struct {
	Browser    string
	Categories map[string]CategoryReport
}

type CategoryReport struct {
	Files     int
	TotalSize int64
}

func Scan(b Browser, verbose bool) (*ScanReport, error) {
	report := &ScanReport{
		Browser:    b.Name,
		Categories: make(map[string]CategoryReport),
	}

	for _, dp := range b.Paths {
		cr := CategoryReport{}

		err := filepath.Walk(dp.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if !info.IsDir() {
				cr.Files++
				cr.TotalSize += info.Size()
				if verbose {
					fmt.Printf("  %s (%s)\n", path, formatBytes(info.Size()))
				}
			}
			return nil
		})

		if err != nil {
			continue
		}

		report.Categories[dp.Category] = cr
	}

	return report, nil
}

func (r *ScanReport) Print() {
	color.Green("[%s]", r.Browser)

	var totalFiles int
	var totalSize int64

	for cat, cr := range r.Categories {
		fmt.Printf("  %-12s %4d files  %s\n", cat, cr.Files, formatBytes(cr.TotalSize))
		totalFiles += cr.Files
		totalSize += cr.TotalSize
	}

	color.White("  %-12s %4d files  %s\n", "total", totalFiles, formatBytes(totalSize))
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
