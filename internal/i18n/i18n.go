// Package i18n provides translation lookup for UI chrome strings.
// Content strings (lesson prose, quizzes, callouts) live in .md files
// and are handled by the content package, not here.
package i18n

import (
	"io/fs"
	"sync"
)

var (
	mu       sync.RWMutex
	catalogs map[string]catalog
	locale   string = "en"
)

// Init loads all locale catalogs from the given filesystem.
// The filesystem should contain directories like en/, es/, etc.,
// each with a ui.toml file.
// Must be called before any other function in this package.
func Init(fsys fs.FS) error {
	loaded, err := loadAllCatalogs(fsys)
	if err != nil {
		return err
	}
	mu.Lock()
	catalogs = loaded
	mu.Unlock()
	return nil
}

// T returns the translated string for key in the current locale.
// Falls back to English if the key is missing in the current locale.
// Returns the key itself if not found in any locale.
func T(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	if cat, ok := catalogs[locale]; ok {
		if val, ok := cat[key]; ok {
			return val
		}
	}

	// Fall back to English.
	if locale != "en" {
		if cat, ok := catalogs["en"]; ok {
			if val, ok := cat[key]; ok {
				return val
			}
		}
	}

	return key
}

// TFmt returns a translated string with interpolated values.
// Keys in the template like {{name}} are replaced with values from args.
func TFmt(key string, args map[string]string) string {
	template := T(key)
	return interpolate(template, args)
}

// TPlural returns the correct plural form for count.
// It looks up key + "_one" or key + "_other" based on count.
// Falls back to key + "_other" if the specific form is missing.
// Additional CLDR categories (zero, two, few, many) can be added
// by extending the suffix lookup without changing this signature.
func TPlural(key string, count int, args map[string]string) string {
	suffix := "_other"
	if count == 1 {
		suffix = "_one"
	}

	mu.RLock()
	defer mu.RUnlock()

	// Try locale-specific plural form.
	pluralKey := key + suffix
	if cat, ok := catalogs[locale]; ok {
		if val, ok := cat[pluralKey]; ok {
			return interpolate(val, args)
		}
	}

	// Fall back to English plural form.
	if locale != "en" {
		if cat, ok := catalogs["en"]; ok {
			if val, ok := cat[pluralKey]; ok {
				return interpolate(val, args)
			}
		}
	}

	// Fall back to _other form.
	if suffix != "_other" {
		otherKey := key + "_other"
		if cat, ok := catalogs[locale]; ok {
			if val, ok := cat[otherKey]; ok {
				return interpolate(val, args)
			}
		}
		if locale != "en" {
			if cat, ok := catalogs["en"]; ok {
				if val, ok := cat[otherKey]; ok {
					return interpolate(val, args)
				}
			}
		}
	}

	return key
}

// SetLocale switches the active locale.
// If the locale is not loaded, falls back to "en".
func SetLocale(loc string) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := catalogs[loc]; ok {
		locale = loc
	} else {
		locale = "en"
	}
}

// Locale returns the current active locale code.
func Locale() string {
	mu.RLock()
	defer mu.RUnlock()
	return locale
}

// Available returns all loaded locale codes.
func Available() []string {
	mu.RLock()
	defer mu.RUnlock()

	locales := make([]string, 0, len(catalogs))
	for loc := range catalogs {
		locales = append(locales, loc)
	}
	return locales
}
