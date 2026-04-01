package content

import (
	"fmt"
	"io/fs"
	"strings"
)

// Resolver maps lesson sections to .md file paths using convention-based
// counters. Each section type (text, callout, quiz) has its own counter
// per lesson. Paths are derived, never explicit in TOML.
type Resolver struct {
	contentFS fs.FS // English content root (content/)
	localeFS  fs.FS // Locale root (locales/)
	locale    string
}

// NewResolver creates a resolver for the given filesystems and locale.
func NewResolver(contentFS, localeFS fs.FS, locale string) *Resolver {
	return &Resolver{
		contentFS: contentFS,
		localeFS:  localeFS,
		locale:    locale,
	}
}

// ResolveSections walks a list of raw section descriptors for a lesson,
// increments per-type counters, and returns the .md file path for each
// section that needs one. Returns ("", false) for sections that don't
// load .md content (diagram, interactive, formula, code).
func ResolveSections(lessonID string) *PathCounter {
	return &PathCounter{lessonID: lessonID}
}

// PathCounter tracks per-type counters for a single lesson.
type PathCounter struct {
	lessonID string
	text     int
	callout  int
	quiz     int
}

// chapterDir extracts the chapter directory from a lesson ID like "ch01.lesson02".
func (pc *PathCounter) chapterDir() string {
	parts := strings.SplitN(pc.lessonID, ".", 2)
	if len(parts) < 1 {
		return ""
	}
	return parts[0]
}

// lessonPrefix extracts the lesson prefix from an ID like "ch01.lesson02".
func (pc *PathCounter) lessonPrefix() string {
	parts := strings.SplitN(pc.lessonID, ".", 2)
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

// NextText returns the path for the next text section.
// e.g., "ch01/lesson01_01.md", "ch01/lesson01_02.md"
func (pc *PathCounter) NextText() string {
	pc.text++
	return fmt.Sprintf("%s/%s_%02d.md", pc.chapterDir(), pc.lessonPrefix(), pc.text)
}

// NextCallout returns the path for the next callout section.
// e.g., "ch01/lesson01_callout01.md"
func (pc *PathCounter) NextCallout() string {
	pc.callout++
	return fmt.Sprintf("%s/%s_callout%02d.md", pc.chapterDir(), pc.lessonPrefix(), pc.callout)
}

// NextQuiz returns the path for the next quiz section.
// e.g., "ch01/lesson01_quiz01.md"
func (pc *PathCounter) NextQuiz() string {
	pc.quiz++
	return fmt.Sprintf("%s/%s_quiz%02d.md", pc.chapterDir(), pc.lessonPrefix(), pc.quiz)
}

// ExerciseTitlePath returns the title .md path for an exercise.
// e.g., "ch01/exercises/ex01_hash_title.md"
func ExerciseTitlePath(chapterID, exerciseID string) string {
	return fmt.Sprintf("%s/exercises/%s_title.md", chapterID, exerciseID)
}

// ExerciseDescPath returns the description .md path for an exercise.
func ExerciseDescPath(chapterID, exerciseID string) string {
	return fmt.Sprintf("%s/exercises/%s_description.md", chapterID, exerciseID)
}

// ReadContent reads a .md file, trying the locale-specific path first,
// falling back to the English content path.
func (r *Resolver) ReadContent(path string) (string, error) {
	// Try locale-specific file first (if not English).
	if r.locale != "en" && r.localeFS != nil {
		localePath := r.locale + "/" + path
		data, err := fs.ReadFile(r.localeFS, localePath)
		if err == nil {
			return string(data), nil
		}
	}

	// Fall back to English content.
	data, err := fs.ReadFile(r.contentFS, path)
	if err != nil {
		return "", fmt.Errorf("content file %s: %w", path, err)
	}
	return string(data), nil
}
