package cleaner

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Browser struct {
	Name     string
	Paths    []DataPath
	Detected bool
}

type DataPath struct {
	Category string
	Path     string
}

func ResolveBrowsers(filter string) []Browser {
	all := detectBrowsers()
	filter = strings.ToLower(filter)

	if filter == "all" {
		return all
	}

	var result []Browser
	for _, b := range all {
		if strings.ToLower(b.Name) == filter {
			result = append(result, b)
		}
	}
	return result
}

func detectBrowsers() []Browser {
	var browsers []Browser

	definitions := browserDefinitions()
	for _, def := range definitions {
		var validPaths []DataPath
		for _, dp := range def.Paths {
			expanded := expandPath(dp.Path)
			if _, err := os.Stat(expanded); err == nil {
				validPaths = append(validPaths, DataPath{
					Category: dp.Category,
					Path:     expanded,
				})
			}
		}
		if len(validPaths) > 0 {
			browsers = append(browsers, Browser{
				Name:     def.Name,
				Paths:    validPaths,
				Detected: true,
			})
		}
	}

	return browsers
}

func browserDefinitions() []Browser {
	switch runtime.GOOS {
	case "linux":
		return linuxBrowsers()
	case "darwin":
		return macosBrowsers()
	case "windows":
		return windowsBrowsers()
	default:
		return nil
	}
}

func linuxBrowsers() []Browser {
	home, _ := os.UserHomeDir()
	return []Browser{
		{
			Name: "Chrome",
			Paths: []DataPath{
				{"cookies", filepath.Join(home, ".config", "google-chrome", "Default", "Cookies")},
				{"cache", filepath.Join(home, ".cache", "google-chrome", "Default", "Cache")},
				{"sessions", filepath.Join(home, ".config", "google-chrome", "Default", "Sessions")},
				{"history", filepath.Join(home, ".config", "google-chrome", "Default", "History")},
			},
		},
		{
			Name: "Firefox",
			Paths: []DataPath{
				{"cookies", filepath.Join(home, ".mozilla", "firefox")},
				{"cache", filepath.Join(home, ".cache", "mozilla", "firefox")},
			},
		},
		{
			Name: "Edge",
			Paths: []DataPath{
				{"cookies", filepath.Join(home, ".config", "microsoft-edge", "Default", "Cookies")},
				{"cache", filepath.Join(home, ".cache", "microsoft-edge", "Default", "Cache")},
			},
		},
	}
}

func macosBrowsers() []Browser {
	home, _ := os.UserHomeDir()
	return []Browser{
		{
			Name: "Chrome",
			Paths: []DataPath{
				{"cookies", filepath.Join(home, "Library", "Application Support", "Google", "Chrome", "Default", "Cookies")},
				{"cache", filepath.Join(home, "Library", "Caches", "Google", "Chrome", "Default", "Cache")},
			},
		},
		{
			Name: "Firefox",
			Paths: []DataPath{
				{"cookies", filepath.Join(home, "Library", "Application Support", "Firefox", "Profiles")},
				{"cache", filepath.Join(home, "Library", "Caches", "Firefox", "Profiles")},
			},
		},
		{
			Name: "Safari",
			Paths: []DataPath{
				{"cookies", filepath.Join(home, "Library", "Cookies")},
				{"cache", filepath.Join(home, "Library", "Caches", "com.apple.Safari")},
			},
		},
	}
}

func windowsBrowsers() []Browser {
	local := os.Getenv("LOCALAPPDATA")
	return []Browser{
		{
			Name: "Chrome",
			Paths: []DataPath{
				{"cookies", filepath.Join(local, "Google", "Chrome", "User Data", "Default", "Cookies")},
				{"cache", filepath.Join(local, "Google", "Chrome", "User Data", "Default", "Cache")},
				{"history", filepath.Join(local, "Google", "Chrome", "User Data", "Default", "History")},
			},
		},
		{
			Name: "Firefox",
			Paths: []DataPath{
				{"cookies", filepath.Join(os.Getenv("APPDATA"), "Mozilla", "Firefox", "Profiles")},
				{"cache", filepath.Join(local, "Mozilla", "Firefox", "Profiles")},
			},
		},
		{
			Name: "Edge",
			Paths: []DataPath{
				{"cookies", filepath.Join(local, "Microsoft", "Edge", "User Data", "Default", "Cookies")},
				{"cache", filepath.Join(local, "Microsoft", "Edge", "User Data", "Default", "Cache")},
			},
		},
	}
}

func expandPath(p string) string {
	if strings.HasPrefix(p, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, p[1:])
	}
	return p
}
