package content

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/BurntSushi/toml"
)

// rawChapter mirrors the TOML structure before type-safe conversion.
type rawChapter struct {
	Chapter struct {
		ID             string `toml:"id"`
		TitleKey       string `toml:"title_key"`
		DescriptionKey string `toml:"description_key"`
	} `toml:"chapter"`
	Lessons []rawLesson `toml:"lessons"`
}

type rawLesson struct {
	ID       string       `toml:"id"`
	TitleKey string       `toml:"title_key"`
	Sections []rawSection `toml:"sections"`
}

type rawSection struct {
	Type       string `toml:"type"`
	Style      string `toml:"style"`
	DiagramID  string `toml:"diagram_id"`
	ExerciseID string `toml:"exercise_id"`
	Formula    string `toml:"formula"`
	Language   string `toml:"language"`
	Correct    []int  `toml:"correct"`
	Content    string `toml:"content"` // For code sections with inline content.
}

// LoadChapter parses a chapter TOML file and resolves all .md content.
// Fails fast on any structural or content error.
func LoadChapter(chapterTOML []byte, chapterFile string, resolver *Resolver) (*Chapter, error) {
	var raw rawChapter
	if err := toml.Unmarshal(chapterTOML, &raw); err != nil {
		return nil, fmt.Errorf("parse %s: %w", chapterFile, err)
	}

	ch := &Chapter{
		ID:       raw.Chapter.ID,
		TitleKey: raw.Chapter.TitleKey,
		DescKey:  raw.Chapter.DescriptionKey,
	}

	for i, rl := range raw.Lessons {
		lesson, err := loadLesson(rl, resolver)
		if err != nil {
			return nil, fmt.Errorf("%s: lesson %d (%s): %w", chapterFile, i, rl.ID, err)
		}
		ch.Lessons = append(ch.Lessons, *lesson)
	}

	return ch, nil
}

func loadLesson(rl rawLesson, resolver *Resolver) (*Lesson, error) {
	if rl.ID == "" {
		return nil, fmt.Errorf("lesson missing id")
	}

	lesson := &Lesson{
		ID:       rl.ID,
		TitleKey: rl.TitleKey,
	}

	counter := ResolveSections(rl.ID)

	for i, rs := range rl.Sections {
		sec, err := loadSection(rs, counter, resolver)
		if err != nil {
			return nil, fmt.Errorf("section %d (type %q): %w", i, rs.Type, err)
		}
		if err := sec.Validate(); err != nil {
			return nil, fmt.Errorf("section %d (type %q): validate: %w", i, rs.Type, err)
		}
		lesson.Sections = append(lesson.Sections, sec)
	}

	return lesson, nil
}

func loadSection(rs rawSection, counter *PathCounter, resolver *Resolver) (Section, error) {
	switch rs.Type {
	case "text":
		return loadTextSection(counter, resolver)
	case "callout":
		return loadCalloutSection(rs, counter, resolver)
	case "quiz":
		return loadQuizSection(rs, counter, resolver)
	case "diagram":
		return &DiagramSection{DiagramID: rs.DiagramID}, nil
	case "interactive":
		return &InteractiveSection{ExerciseID: rs.ExerciseID}, nil
	case "formula":
		return &FormulaSection{Formula: rs.Formula}, nil
	case "code":
		return &CodeSection{Content: rs.Content, Language: rs.Language}, nil
	default:
		return nil, fmt.Errorf("unknown section type %q", rs.Type)
	}
}

func loadTextSection(counter *PathCounter, resolver *Resolver) (Section, error) {
	path := counter.NextText()
	content, err := resolver.ReadContent(path)
	if err != nil {
		return nil, err
	}
	return &TextSection{Content: content}, nil
}

func loadCalloutSection(rs rawSection, counter *PathCounter, resolver *Resolver) (Section, error) {
	path := counter.NextCallout()
	content, err := resolver.ReadContent(path)
	if err != nil {
		return nil, err
	}
	return &CalloutSection{Content: content, Style: rs.Style}, nil
}

func loadQuizSection(rs rawSection, counter *PathCounter, resolver *Resolver) (Section, error) {
	path := counter.NextQuiz()
	content, err := resolver.ReadContent(path)
	if err != nil {
		return nil, err
	}

	question, options, err := ParseQuizMD(content)
	if err != nil {
		return nil, fmt.Errorf("quiz %s: %w", path, err)
	}

	return &QuizSection{
		Question: question,
		Options:  options,
		Correct:  rs.Correct,
	}, nil
}

// ParseQuizMD extracts the question and options from a quiz .md file.
// First non-empty line is the question. Lines starting with "- " are options.
func ParseQuizMD(content string) (string, []string, error) {
	lines := strings.Split(content, "\n")

	var question string
	var options []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			options = append(options, strings.TrimPrefix(trimmed, "- "))
		} else if question == "" {
			question = trimmed
		}
	}

	if question == "" {
		return "", nil, fmt.Errorf("no question found (expected first non-empty line)")
	}
	if len(options) < 2 {
		return "", nil, fmt.Errorf("need at least 2 options (lines starting with \"- \"), got %d", len(options))
	}

	return question, options, nil
}

// LoadChapterFromFS loads a chapter TOML file from a filesystem.
func LoadChapterFromFS(fsys fs.FS, path string, resolver *Resolver) (*Chapter, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return LoadChapter(data, path, resolver)
}
