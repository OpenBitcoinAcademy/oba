package i18n

import (
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

// catalog holds flattened key-value pairs for a single locale.
type catalog map[string]string

// loadCatalog reads a ui.toml file and flattens it into dot-separated keys.
// For example, [nav] home = "Home" becomes "nav.home" -> "Home".
func loadCatalog(fsys fs.FS, path string) (catalog, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var raw map[string]any
	if err := toml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	cat := make(catalog)
	flatten(cat, "", raw)
	return cat, nil
}

// flatten recursively walks a nested map and produces dot-separated keys.
// Non-string, non-map values are logged as warnings (fail-fast at startup
// would be better, but changing the signature is deferred).
func flatten(cat catalog, prefix string, m map[string]any) {
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch val := v.(type) {
		case string:
			cat[key] = val
		case map[string]any:
			flatten(cat, key, val)
		default:
			log.Printf("i18n: unexpected value type for key %q: %T", key, v)
		}
	}
}

// loadAllCatalogs walks the filesystem looking for {locale}/ui.toml files
// and loads each into a catalog.
func loadAllCatalogs(fsys fs.FS) (map[string]catalog, error) {
	catalogs := make(map[string]catalog)

	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, fmt.Errorf("read locale directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		locale := entry.Name()
		path := locale + "/ui.toml"

		// Check if ui.toml exists in this locale dir.
		if _, err := fs.Stat(fsys, path); err != nil {
			continue
		}

		cat, err := loadCatalog(fsys, path)
		if err != nil {
			return nil, fmt.Errorf("load locale %s: %w", locale, err)
		}
		catalogs[locale] = cat
	}

	if _, ok := catalogs["en"]; !ok {
		return nil, fmt.Errorf("english locale (en/ui.toml) is required but missing")
	}

	return catalogs, nil
}

// interpolate replaces {{key}} placeholders with values from args.
func interpolate(template string, args map[string]string) string {
	result := template
	for k, v := range args {
		result = strings.ReplaceAll(result, "{{"+k+"}}", v)
	}
	return result
}
