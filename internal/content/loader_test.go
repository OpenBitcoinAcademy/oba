package content

import (
	"strings"
	"testing"
	"testing/fstest"
)

func testContentFS() fstest.MapFS {
	return fstest.MapFS{
		"ch01/lesson01_01.md":        {Data: []byte("What is money?")},
		"ch01/lesson01_callout01.md": {Data: []byte("Inflation warning.")},
		"ch01/lesson01_02.md":        {Data: []byte("The trust problem.")},
		"ch01/lesson01_quiz01.md":    {Data: []byte("What does Bitcoin solve?\n\n- Speed\n- Trust\n- Email\n")},
	}
}

var testTOML = `
[chapter]
id = "ch01"
title_key = "chapters.ch01_title"
description_key = "chapters.ch01_description"

[[lessons]]
id = "ch01.lesson01"
title_key = "lessons.ch01_lesson01_title"

  [[lessons.sections]]
  type = "text"

  [[lessons.sections]]
  type = "callout"
  style = "info"

  [[lessons.sections]]
  type = "text"

  [[lessons.sections]]
  type = "diagram"
  diagram_id = "centralized_vs_p2p"

  [[lessons.sections]]
  type = "quiz"
  correct = [1]
`

func TestLoadChapter_Basic(t *testing.T) {
	cfs := testContentFS()
	resolver := NewResolver(cfs, nil, "en")

	ch, err := LoadChapter([]byte(testTOML), "ch01.toml", resolver)
	if err != nil {
		t.Fatal(err)
	}

	if ch.ID != "ch01" {
		t.Errorf("chapter ID = %q", ch.ID)
	}
	if len(ch.Lessons) != 1 {
		t.Fatalf("lessons = %d, want 1", len(ch.Lessons))
	}

	lesson := ch.Lessons[0]
	if lesson.ID != "ch01.lesson01" {
		t.Errorf("lesson ID = %q", lesson.ID)
	}
	if len(lesson.Sections) != 5 {
		t.Fatalf("sections = %d, want 5", len(lesson.Sections))
	}

	// Verify section types.
	types := []string{"text", "callout", "text", "diagram", "quiz"}
	for i, sec := range lesson.Sections {
		if sec.SectionType() != types[i] {
			t.Errorf("section %d type = %q, want %q", i, sec.SectionType(), types[i])
		}
	}
}

func TestLoadChapter_TextContent(t *testing.T) {
	cfs := testContentFS()
	resolver := NewResolver(cfs, nil, "en")

	ch, err := LoadChapter([]byte(testTOML), "ch01.toml", resolver)
	if err != nil {
		t.Fatal(err)
	}

	text1 := ch.Lessons[0].Sections[0].(*TextSection)
	if text1.Content != "What is money?" {
		t.Errorf("text1 = %q", text1.Content)
	}
	text2 := ch.Lessons[0].Sections[2].(*TextSection)
	if text2.Content != "The trust problem." {
		t.Errorf("text2 = %q", text2.Content)
	}
}

func TestLoadChapter_CalloutContent(t *testing.T) {
	cfs := testContentFS()
	resolver := NewResolver(cfs, nil, "en")

	ch, err := LoadChapter([]byte(testTOML), "ch01.toml", resolver)
	if err != nil {
		t.Fatal(err)
	}

	callout := ch.Lessons[0].Sections[1].(*CalloutSection)
	if callout.Content != "Inflation warning." {
		t.Errorf("callout = %q", callout.Content)
	}
	if callout.Style != "info" {
		t.Errorf("style = %q", callout.Style)
	}
}

func TestLoadChapter_QuizParsing(t *testing.T) {
	cfs := testContentFS()
	resolver := NewResolver(cfs, nil, "en")

	ch, err := LoadChapter([]byte(testTOML), "ch01.toml", resolver)
	if err != nil {
		t.Fatal(err)
	}

	quiz := ch.Lessons[0].Sections[4].(*QuizSection)
	if quiz.Question != "What does Bitcoin solve?" {
		t.Errorf("question = %q", quiz.Question)
	}
	if len(quiz.Options) != 3 {
		t.Fatalf("options = %d, want 3", len(quiz.Options))
	}
	if quiz.Options[1] != "Trust" {
		t.Errorf("option[1] = %q", quiz.Options[1])
	}
	if len(quiz.Correct) != 1 || quiz.Correct[0] != 1 {
		t.Errorf("correct = %v", quiz.Correct)
	}
}

func TestLoadChapter_DiagramSection(t *testing.T) {
	cfs := testContentFS()
	resolver := NewResolver(cfs, nil, "en")

	ch, err := LoadChapter([]byte(testTOML), "ch01.toml", resolver)
	if err != nil {
		t.Fatal(err)
	}

	diag := ch.Lessons[0].Sections[3].(*DiagramSection)
	if diag.DiagramID != "centralized_vs_p2p" {
		t.Errorf("diagram_id = %q", diag.DiagramID)
	}
}

func TestLoadChapter_MissingContentFile(t *testing.T) {
	emptyFS := fstest.MapFS{}
	resolver := NewResolver(emptyFS, nil, "en")

	_, err := LoadChapter([]byte(testTOML), "ch01.toml", resolver)
	if err == nil {
		t.Fatal("expected error for missing .md file")
	}
}

func TestLoadChapter_UnknownSectionType(t *testing.T) {
	badTOML := `
[chapter]
id = "ch01"
title_key = "t"
description_key = "d"

[[lessons]]
id = "ch01.lesson01"
title_key = "t"

  [[lessons.sections]]
  type = "bogus"
`
	resolver := NewResolver(fstest.MapFS{}, nil, "en")
	_, err := LoadChapter([]byte(badTOML), "ch01.toml", resolver)
	if err == nil {
		t.Fatal("expected error for unknown type")
	}
	if !strings.Contains(err.Error(), "bogus") {
		t.Errorf("error should mention 'bogus': %v", err)
	}
}

func TestLoadChapter_MalformedTOML(t *testing.T) {
	resolver := NewResolver(fstest.MapFS{}, nil, "en")
	_, err := LoadChapter([]byte("not valid {{{ toml"), "bad.toml", resolver)
	if err == nil {
		t.Fatal("expected error for malformed TOML")
	}
}

func TestLoadChapter_MissingLessonID(t *testing.T) {
	badTOML := `
[chapter]
id = "ch01"
title_key = "t"
description_key = "d"

[[lessons]]
title_key = "t"
`
	resolver := NewResolver(fstest.MapFS{}, nil, "en")
	_, err := LoadChapter([]byte(badTOML), "ch01.toml", resolver)
	if err == nil {
		t.Fatal("expected error for missing lesson ID")
	}
}

// --- Quiz .md parsing ---

func TestParseQuizMD_Basic(t *testing.T) {
	md := "What is 2+2?\n\n- 3\n- 4\n- 5\n"
	q, opts, err := ParseQuizMD(md)
	if err != nil {
		t.Fatal(err)
	}
	if q != "What is 2+2?" {
		t.Errorf("question = %q", q)
	}
	if len(opts) != 3 {
		t.Fatalf("options = %d", len(opts))
	}
	if opts[1] != "4" {
		t.Errorf("option[1] = %q", opts[1])
	}
}

func TestParseQuizMD_NoQuestion(t *testing.T) {
	md := "- A\n- B\n"
	_, _, err := ParseQuizMD(md)
	if err == nil {
		t.Fatal("expected error: no question")
	}
}

func TestParseQuizMD_TooFewOptions(t *testing.T) {
	md := "Question?\n\n- Only one\n"
	_, _, err := ParseQuizMD(md)
	if err == nil {
		t.Fatal("expected error: too few options")
	}
}

func TestParseQuizMD_EmptyContent(t *testing.T) {
	_, _, err := ParseQuizMD("")
	if err == nil {
		t.Fatal("expected error: empty content")
	}
}

func TestParseQuizMD_BlankLinesBetween(t *testing.T) {
	md := "\n\nQuestion here\n\n\n- A\n\n- B\n\n- C\n\n"
	q, opts, err := ParseQuizMD(md)
	if err != nil {
		t.Fatal(err)
	}
	if q != "Question here" {
		t.Errorf("question = %q", q)
	}
	if len(opts) != 3 {
		t.Errorf("options = %d", len(opts))
	}
}
