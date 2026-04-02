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
	ScreenHome    Screen = iota // Chapter selector
	ScreenChapter               // Lesson list for a selected chapter
	ScreenLesson                // Lesson content viewer
	ScreenSettings
)

// InvalidateFunc is called after state changes to trigger a redraw.
type InvalidateFunc func()

// State holds the application-wide state shared across screens.
type State struct {
	Theme     *theme.Theme
	Chapters  []*content.Chapter                 // all loaded chapters
	Chapter   *content.Chapter                   // current chapter (pointer into Chapters)
	Exercises map[string]*content.ExerciseConfig // exercise ID -> config (all chapters)
	Progress  *content.Progress

	CurrentScreen  Screen
	CurrentChapter int  // index into Chapters
	CurrentLesson  int  // index into Chapter.Lessons
	LessonDirty    bool // true when lesson/chapter changed, screens should reset

	ContentFS fs.FS
	LocaleFS  fs.FS

	ProgressPath string
	Invalidate   InvalidateFunc
}

// NewState creates the initial application state with default fonts.
func NewState(chapters []*content.Chapter, progress *content.Progress, progressPath string) *State {
	return NewStateWithFonts(chapters, progress, progressPath, nil, nil)
}

// NewStateWithFonts creates the initial state with custom font data.
func NewStateWithFonts(chapters []*content.Chapter, progress *content.Progress, progressPath string, notoSans, jetBrainsMono []byte) *State {
	var ch *content.Chapter
	if len(chapters) > 0 {
		ch = chapters[0]
	}
	return &State{
		Theme:         theme.NewWithFonts(theme.DetectSystemMode(), notoSans, jetBrainsMono),
		Chapters:      chapters,
		Chapter:       ch,
		Progress:      progress,
		CurrentScreen: ScreenHome,
		ProgressPath:  progressPath,
		Invalidate:    func() {},
	}
}

// NavigateToChapter selects a chapter and shows its lesson list.
func (s *State) NavigateToChapter(index int) {
	if index >= 0 && index < len(s.Chapters) {
		s.CurrentChapter = index
		s.Chapter = s.Chapters[index]
		s.CurrentScreen = ScreenChapter
		s.LessonDirty = true
		s.Invalidate()
	}
}

// NavigateToLesson switches to the lesson view and signals a reset.
func (s *State) NavigateToLesson(index int) {
	if s.Chapter != nil && index >= 0 && index < len(s.Chapter.Lessons) {
		s.CurrentLesson = index
		s.CurrentScreen = ScreenLesson
		s.LessonDirty = true
		s.Invalidate()
	}
}

// NavigateHome returns to the chapter selector.
func (s *State) NavigateHome() {
	s.CurrentScreen = ScreenHome
	s.Invalidate()
}

// NavigateChapterList returns to the lesson list for the current chapter.
func (s *State) NavigateChapterList() {
	s.CurrentScreen = ScreenChapter
	s.Invalidate()
}

// NavigateSettings opens the settings screen.
func (s *State) NavigateSettings() {
	s.CurrentScreen = ScreenSettings
	s.Invalidate()
}

// SetLocale switches the UI language and reloads all chapter content.
func (s *State) SetLocale(locale string) {
	i18n.SetLocale(locale)
	s.Progress.Locale = locale

	if s.ContentFS != nil {
		resolver := content.NewResolver(s.ContentFS, s.LocaleFS, locale)
		chapters, err := content.LoadAllChapters(s.ContentFS, "chapters.toml", resolver)
		if err != nil {
			log.Printf("reload content for locale %s: %v (keeping current)", locale, err)
		} else {
			s.Chapters = chapters
			if s.CurrentChapter < len(chapters) {
				s.Chapter = chapters[s.CurrentChapter]
			}
			s.LessonDirty = true
		}
	}

	s.saveProgress()
	s.Invalidate()
}

// CompleteLesson marks the current lesson as done, saves, and invalidates.
func (s *State) CompleteLesson() {
	if s.Chapter != nil && s.CurrentLesson >= 0 && s.CurrentLesson < len(s.Chapter.Lessons) {
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
