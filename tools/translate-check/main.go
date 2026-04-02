// Command translate-check validates locale completeness and detects stale
// translations. It walks chapter TOML files, resolves expected .md paths
// for English, then checks each locale for missing files.
//
// Exit codes:
//
//	0 — all English files present, locale warnings only
//	1 — missing English .md files (content is broken)
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type chaptersFile struct {
	Chapters []struct {
		ID       string `toml:"id"`
		TOMLFile string `toml:"toml_file"`
	} `toml:"chapters"`
}

type chapterFile struct {
	Lessons []struct {
		ID       string `toml:"id"`
		Sections []struct {
			Type string `toml:"type"`
		} `toml:"sections"`
	} `toml:"lessons"`
}

func main() {
	contentDir := "content"
	localesDir := "locales"

	// Load chapter index.
	indexData, err := os.ReadFile(filepath.Join(contentDir, "chapters.toml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "read chapters.toml: %v\n", err)
		os.Exit(1)
	}

	var index chaptersFile
	if err := toml.Unmarshal(indexData, &index); err != nil {
		fmt.Fprintf(os.Stderr, "parse chapters.toml: %v\n", err)
		os.Exit(1)
	}

	var englishPaths []string
	errors := 0
	warnings := 0

	for _, ch := range index.Chapters {
		chData, err := os.ReadFile(filepath.Join(contentDir, ch.TOMLFile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: read %s: %v\n", ch.TOMLFile, err)
			errors++
			continue
		}

		var chapter chapterFile
		if err := toml.Unmarshal(chData, &chapter); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: parse %s: %v\n", ch.TOMLFile, err)
			errors++
			continue
		}

		for _, lesson := range chapter.Lessons {
			paths := resolvePaths(lesson.ID, lesson.Sections)
			englishPaths = append(englishPaths, paths...)
		}
	}

	// Check all English files exist.
	fmt.Printf("Checking %d English .md files...\n", len(englishPaths))
	for _, p := range englishPaths {
		full := filepath.Join(contentDir, p)
		if _, err := os.Stat(full); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: missing English file: %s\n", full)
			errors++
		}
	}

	// Find all locales (directories under locales/ that aren't "en").
	locales, err := findLocales(localesDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read locales dir: %v\n", err)
		os.Exit(1)
	}

	// Check each locale for missing and stale translations.
	for _, locale := range locales {
		missing := 0
		stale := 0
		for _, p := range englishPaths {
			enFull := filepath.Join(contentDir, p)
			locFull := filepath.Join(localesDir, locale, p)

			locStat, err := os.Stat(locFull)
			if err != nil {
				missing++
				continue
			}

			// Stale check: compare content hashes.
			enHash := fileHash(enFull)
			locHash := fileHash(locFull)
			if enHash == locHash {
				// Identical content means it was copied, not translated.
				// This is a heuristic: skip for now.
				_ = locStat
			}

			// Check if English file is newer than translation.
			enStat, _ := os.Stat(enFull)
			if enStat != nil && locStat != nil && enStat.ModTime().After(locStat.ModTime()) {
				stale++
			}
		}

		total := len(englishPaths)
		translated := total - missing
		fmt.Printf("  %s: %d/%d translated", locale, translated, total)
		if stale > 0 {
			fmt.Printf(", %d potentially stale", stale)
		}
		if missing > 0 {
			fmt.Printf(", %d missing", missing)
			warnings += missing
		}
		fmt.Println()
	}

	if errors > 0 {
		fmt.Fprintf(os.Stderr, "\n%d errors found. English content is incomplete.\n", errors)
		os.Exit(1)
	}

	if warnings > 0 {
		fmt.Printf("\n%d translation warnings (missing locale files).\n", warnings)
	} else {
		fmt.Println("\nAll checks passed.")
	}
}

func resolvePaths(lessonID string, sections []struct {
	Type string `toml:"type"`
}) []string {
	parts := strings.SplitN(lessonID, ".", 2)
	if len(parts) < 2 {
		return nil
	}
	chDir := parts[0]
	prefix := parts[1]

	var paths []string
	textN, calloutN, quizN := 0, 0, 0

	for _, sec := range sections {
		switch sec.Type {
		case "text":
			textN++
			paths = append(paths, fmt.Sprintf("%s/%s_%02d.md", chDir, prefix, textN))
		case "callout":
			calloutN++
			paths = append(paths, fmt.Sprintf("%s/%s_callout%02d.md", chDir, prefix, calloutN))
		case "quiz":
			quizN++
			paths = append(paths, fmt.Sprintf("%s/%s_quiz%02d.md", chDir, prefix, quizN))
		}
	}
	return paths
}

func findLocales(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var locales []string
	for _, e := range entries {
		if e.IsDir() && e.Name() != "en" {
			// Check it has a ui.toml.
			if _, err := os.Stat(filepath.Join(dir, e.Name(), "ui.toml")); err == nil {
				locales = append(locales, e.Name())
			}
		}
	}
	return locales, nil
}

func fileHash(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	h := sha256.Sum256(data)
	return fmt.Sprintf("%x", h[:8])
}

// Ensure fs import is used.
