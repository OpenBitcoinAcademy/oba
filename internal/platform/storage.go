// Package platform provides platform-specific abstractions for storage
// and system services. Desktop uses XDG directories on Linux and
// AppData on Windows.
package platform

import (
	"os"
	"path/filepath"
	"runtime"
)

// DataDir returns the platform-specific directory for persistent app data.
// Creates the directory if it does not exist.
func DataDir() (string, error) {
	dir := dataDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return dir, nil
}

// ProgressPath returns the full path to the progress JSON file.
func ProgressPath() (string, error) {
	dir, err := DataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "progress.json"), nil
}

func dataDir() string {
	switch runtime.GOOS {
	case "windows":
		base := os.Getenv("APPDATA")
		if base == "" {
			base = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
		}
		return filepath.Join(base, "OpenBitcoinAcademy")
	case "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "Library", "Application Support", "OpenBitcoinAcademy")
	case "android":
		// Android app-private storage provided by Gio.
		dir, _ := os.UserCacheDir()
		return filepath.Join(filepath.Dir(dir), "files")
	default:
		// Linux/BSD: follow XDG Base Directory spec.
		// Also handles Android (runtime.GOOS == "linux") where
		// UserHomeDir may not exist. Fall back to cache dir parent.
		base := os.Getenv("XDG_DATA_HOME")
		if base == "" {
			home, err := os.UserHomeDir()
			if err != nil || home == "" {
				// Android fallback: use cache dir's parent.
				if cacheDir, err := os.UserCacheDir(); err == nil {
					return filepath.Join(filepath.Dir(cacheDir), "files")
				}
				return filepath.Join(os.TempDir(), "open-bitcoin-academy")
			}
			base = filepath.Join(home, ".local", "share")
		}
		return filepath.Join(base, "open-bitcoin-academy")
	}
}
