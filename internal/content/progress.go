package content

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Progress tracks completed lessons and exercises.
type Progress struct {
	Version   int                  `json:"version"`
	Locale    string               `json:"locale"`
	Completed map[string]time.Time `json:"completed"`
	Exercises map[string]time.Time `json:"exercises"`
}

const progressVersion = 1

// NewProgress creates a fresh progress state.
func NewProgress() *Progress {
	return &Progress{
		Version:   progressVersion,
		Locale:    "en",
		Completed: make(map[string]time.Time),
		Exercises: make(map[string]time.Time),
	}
}

// MarkLessonComplete records a lesson as completed.
func (p *Progress) MarkLessonComplete(lessonID string) {
	p.Completed[lessonID] = time.Now()
}

// MarkExerciseComplete records an exercise as completed.
func (p *Progress) MarkExerciseComplete(exerciseID string) {
	p.Exercises[exerciseID] = time.Now()
}

// IsLessonComplete returns true if the lesson has been completed.
func (p *Progress) IsLessonComplete(lessonID string) bool {
	_, ok := p.Completed[lessonID]
	return ok
}

// IsExerciseComplete returns true if the exercise has been completed.
func (p *Progress) IsExerciseComplete(exerciseID string) bool {
	_, ok := p.Exercises[exerciseID]
	return ok
}

// CompletedCount returns how many lessons have been completed.
func (p *Progress) CompletedCount() int {
	return len(p.Completed)
}

// Reset clears all progress data.
func (p *Progress) Reset() {
	p.Completed = make(map[string]time.Time)
	p.Exercises = make(map[string]time.Time)
}

// LoadProgress reads progress from a JSON file. Returns fresh progress
// if the file is missing, corrupted, or from a newer schema version.
// Never returns an error: bad data resets silently.
func LoadProgress(path string) *Progress {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return NewProgress()
		}
		log.Printf("progress: read %s: %v (resetting)", path, err)
		return NewProgress()
	}

	var p Progress
	if err := json.Unmarshal(data, &p); err != nil {
		log.Printf("progress: parse %s: %v (resetting)", path, err)
		return NewProgress()
	}

	if p.Version > progressVersion {
		log.Printf("progress: version %d > %d (resetting)", p.Version, progressVersion)
		return NewProgress()
	}

	if p.Completed == nil {
		p.Completed = make(map[string]time.Time)
	}
	if p.Exercises == nil {
		p.Exercises = make(map[string]time.Time)
	}

	return &p
}

// SaveProgress writes progress to a JSON file, creating parent
// directories as needed.
func SaveProgress(path string, p *Progress) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create progress dir: %w", err)
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal progress: %w", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write progress: %w", err)
	}
	return nil
}

// LoadProgressFS reads progress from an fs.FS (for testing).
func LoadProgressFS(fsys fs.FS, path string) *Progress {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return NewProgress()
	}

	var p Progress
	if err := json.Unmarshal(data, &p); err != nil {
		return NewProgress()
	}

	if p.Version > progressVersion {
		return NewProgress()
	}

	if p.Completed == nil {
		p.Completed = make(map[string]time.Time)
	}
	if p.Exercises == nil {
		p.Exercises = make(map[string]time.Time)
	}

	return &p
}
