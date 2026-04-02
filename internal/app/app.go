// Package app manages top-level application state and screen navigation.
package app

import (
	"io/fs"
	"log"

	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
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
	Theme     *theme.Theme
	Chapter   *content.Chapter
	Exercises map[string]*content.ExerciseConfig // exercise ID -> config
	Progress  *content.Progress

	CurrentScreen Screen
	CurrentLesson int  // index into Chapter.Lessons
	LessonDirty   bool // true when lesson index changed, screens should reset

	// Embedded filesystems for content reload on locale switch.
	ContentFS fs.FS
	LocaleFS  fs.FS

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
		Invalidate:    func() {},
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

// SetLocale switches the UI language and reloads lesson content.
func (s *State) SetLocale(locale string) {
	i18n.SetLocale(locale)
	s.Progress.Locale = locale

	// Reload chapter content with the new locale's resolver.
	if s.ContentFS != nil {
		resolver := content.NewResolver(s.ContentFS, s.LocaleFS, locale)
		ch, err := content.LoadChapterFromFS(s.ContentFS, "ch01/ch01.toml", resolver)
		if err != nil {
			log.Printf("reload content for locale %s: %v (keeping current)", locale, err)
		} else {
			s.Chapter = ch
			s.LessonDirty = true
		}
	}

	s.saveProgress()
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
