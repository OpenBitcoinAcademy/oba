// Package app manages top-level application state and screen navigation.
package app

import (
	"log"

	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Screen identifies the current view.
type Screen int

const (
	ScreenHome Screen = iota
	ScreenLesson
	ScreenSettings
)

// InvalidateFunc is called after state changes to trigger a redraw.
type InvalidateFunc func()

// State holds the application-wide state shared across screens.
type State struct {
	Theme    *theme.Theme
	Chapter  *content.Chapter
	Progress *content.Progress

	CurrentScreen Screen
	CurrentLesson int  // index into Chapter.Lessons
	LessonDirty   bool // true when lesson index changed, screens should reset

	ProgressPath string // file path for saving progress
	Invalidate   InvalidateFunc
}

// NewState creates the initial application state.
func NewState(ch *content.Chapter, progress *content.Progress, progressPath string) *State {
	return &State{
		Theme:         theme.New(),
		Chapter:       ch,
		Progress:      progress,
		CurrentScreen: ScreenHome,
		ProgressPath:  progressPath,
		Invalidate:    func() {}, // no-op until wired
	}
}

// NavigateToLesson switches to the lesson view and signals a reset.
func (s *State) NavigateToLesson(index int) {
	if index >= 0 && index < len(s.Chapter.Lessons) {
		s.CurrentLesson = index
		s.CurrentScreen = ScreenLesson
		s.LessonDirty = true
		s.Invalidate()
	}
}

// NavigateHome returns to the lesson list.
func (s *State) NavigateHome() {
	s.CurrentScreen = ScreenHome
	s.Invalidate()
}

// NavigateSettings opens the settings screen.
func (s *State) NavigateSettings() {
	s.CurrentScreen = ScreenSettings
	s.Invalidate()
}

// CompleteLesson marks the current lesson as done, saves, and invalidates.
func (s *State) CompleteLesson() {
	if s.CurrentLesson >= 0 && s.CurrentLesson < len(s.Chapter.Lessons) {
		lesson := s.Chapter.Lessons[s.CurrentLesson]
		s.Progress.MarkLessonComplete(lesson.ID)
		s.saveProgress()
		s.Invalidate()
	}
}

// SetThemeMode switches between light and dark mode.
func (s *State) SetThemeMode(mode theme.Mode) {
	s.Theme.SetMode(mode)
	s.Invalidate()
}

// CompleteExercise marks an exercise as done and saves.
func (s *State) CompleteExercise(exerciseID string) {
	s.Progress.MarkExerciseComplete(exerciseID)
	s.saveProgress()
}

func (s *State) saveProgress() {
	if s.ProgressPath == "" {
		return
	}
	if err := content.SaveProgress(s.ProgressPath, s.Progress); err != nil {
		log.Printf("save progress: %v", err)
	}
}
