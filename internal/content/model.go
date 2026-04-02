// Package content handles TOML+Markdown content loading for lessons,
// quizzes, callouts, and exercises. Structure lives in TOML. Prose
// lives in Markdown files resolved by convention from section position.
package content

import "fmt"

// Chapter is the top-level content container.
type Chapter struct {
	ID       string
	TitleKey string
	DescKey  string
	Lessons  []Lesson
}

// Lesson holds an ordered sequence of sections.
type Lesson struct {
	ID       string
	TitleKey string
	Sections []Section
}

// Section is implemented by every section type. Each concrete type
// carries only its own fields, enforced at compile time.
type Section interface {
	SectionType() string
	Validate() error
}

// TextSection holds prose loaded from a .md file.
type TextSection struct {
	Content string
}

func (s *TextSection) SectionType() string { return "text" }

func (s *TextSection) Validate() error {
	if s.Content == "" {
		return fmt.Errorf("text section has empty content")
	}
	return nil
}

// CalloutSection holds a highlighted text block loaded from a .md file.
type CalloutSection struct {
	Content string
	Style   string // "info", "warning", "tip"
}

func (s *CalloutSection) SectionType() string { return "callout" }

func (s *CalloutSection) Validate() error {
	if s.Content == "" {
		return fmt.Errorf("callout section has empty content")
	}
	switch s.Style {
	case "info", "warning", "tip":
		return nil
	case "":
		return fmt.Errorf("callout section missing style")
	default:
		return fmt.Errorf("callout section has unknown style %q", s.Style)
	}
}

// DiagramSection references a registered Gio-native diagram by ID.
type DiagramSection struct {
	DiagramID string
}

func (s *DiagramSection) SectionType() string { return "diagram" }

// ValidateDiagramID is called during Validate to check that a diagram ID
// exists in the registry. Set this before loading content. If nil, only
// checks that the ID is non-empty.
var ValidateDiagramID func(id string) error

func (s *DiagramSection) Validate() error {
	if s.DiagramID == "" {
		return fmt.Errorf("diagram section missing diagram_id")
	}
	if ValidateDiagramID != nil {
		return ValidateDiagramID(s.DiagramID)
	}
	return nil
}

// QuizSection holds a question, options, and correct answer indices.
// Question and Options are parsed from a .md file. Correct indices
// come from TOML.
type QuizSection struct {
	Question string
	Options  []string
	Correct  []int
}

func (s *QuizSection) SectionType() string { return "quiz" }

func (s *QuizSection) Validate() error {
	if s.Question == "" {
		return fmt.Errorf("quiz section has empty question")
	}
	if len(s.Options) < 2 {
		return fmt.Errorf("quiz section needs at least 2 options, got %d", len(s.Options))
	}
	if len(s.Correct) < 1 {
		return fmt.Errorf("quiz section needs at least 1 correct answer")
	}
	for _, idx := range s.Correct {
		if idx < 0 || idx >= len(s.Options) {
			return fmt.Errorf("quiz correct index %d out of range [0, %d)", idx, len(s.Options))
		}
	}
	return nil
}

// InteractiveSection references an exercise by ID.
type InteractiveSection struct {
	ExerciseID string
}

func (s *InteractiveSection) SectionType() string { return "interactive" }

func (s *InteractiveSection) Validate() error {
	if s.ExerciseID == "" {
		return fmt.Errorf("interactive section missing exercise_id")
	}
	return nil
}

// FormulaSection holds a LaTeX subset string for block-level math.
type FormulaSection struct {
	Formula string
}

func (s *FormulaSection) SectionType() string { return "formula" }

func (s *FormulaSection) Validate() error {
	if s.Formula == "" {
		return fmt.Errorf("formula section has empty formula")
	}
	return nil
}

// CodeSection holds syntax-highlighted code for display.
type CodeSection struct {
	Content  string
	Language string
}

func (s *CodeSection) SectionType() string { return "code" }

func (s *CodeSection) Validate() error {
	if s.Content == "" {
		return fmt.Errorf("code section has empty content")
	}
	return nil
}
