package cleanerpackage

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
)
cleaner

import (
"fmt"
"os"
"path/filepath"

"github.com/fatih/color"
)

type CleanResult struct {
	Browser      string
	FilesRemoved int
	BytesFreed   int64
	Errors       int
}

func Clean(b Browser, dryRun bool, verbose bool) (*CleanResult, error) {
	result := &CleanResult{Browser: b.Name}

	for _, dp := range b.Paths {
		err := filepath.Walk(dp.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}

			if dryRun {
				if verbose {
					color.Yellow("  [dry-run] would delete: %s (%s)", path, formatBytes(info.Size()))
				}
				result.FilesRemoved++
				result.BytesFreed += info.Size()
				return nil
			}

			if err := os.Remove(path); err != nil {
				result.Errors++
				if verbose {
					color.Red("  failed: %s — %v", path, err)
				}
				return nil
			}

			result.FilesRemoved++
			result.BytesFreed += info.Size()

			if verbose {
				color.Green("  deleted: %s (%s)", path, formatBytes(info.Size()))
			}

			return nil
		})

		if err != nil {
			result.Errors++
		}
	}

	removeEmptyDirs(b)

	return result, nil
}

func removeEmptyDirs(b Browser) {
	for _, dp := range b.Paths {
		info, err := os.Stat(dp.Path)
		if err != nil || !info.IsDir() {
			continue
		}

		filepath.Walk(dp.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil || !info.IsDir() || path == dp.Path {
				return nil
			}

			entries, err := os.ReadDir(path)
			if err == nil && len(entries) == 0 {
				os.Remove(path)
			}
			return nil
		})
	}
}
