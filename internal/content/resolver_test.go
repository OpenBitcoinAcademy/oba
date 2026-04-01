package content

import (
	"testing"
	"testing/fstest"
)

func TestPathCounter_Text(t *testing.T) {
	pc := ResolveSections("ch01.lesson01")

	tests := []struct {
		want string
	}{
		{"ch01/lesson01_01.md"},
		{"ch01/lesson01_02.md"},
		{"ch01/lesson01_03.md"},
	}
	for i, tt := range tests {
		got := pc.NextText()
		if got != tt.want {
			t.Errorf("text %d: got %q, want %q", i+1, got, tt.want)
		}
	}
}

func TestPathCounter_Callout(t *testing.T) {
	pc := ResolveSections("ch01.lesson02")
	got := pc.NextCallout()
	want := "ch01/lesson02_callout01.md"
	if got != want {
		t.Errorf("callout: got %q, want %q", got, want)
	}
}

func TestPathCounter_Quiz(t *testing.T) {
	pc := ResolveSections("ch01.lesson01")
	got := pc.NextQuiz()
	want := "ch01/lesson01_quiz01.md"
	if got != want {
		t.Errorf("quiz: got %q, want %q", got, want)
	}
}

func TestPathCounter_IndependentCounters(t *testing.T) {
	// Verify that text, callout, and quiz counters are independent.
	pc := ResolveSections("ch01.lesson01")
	pc.NextText()    // text 1
	pc.NextCallout() // callout 1
	pc.NextText()    // text 2
	pc.NextQuiz()    // quiz 1
	got := pc.NextText()
	want := "ch01/lesson01_03.md"
	if got != want {
		t.Errorf("text after interleaved types: got %q, want %q", got, want)
	}
}

func TestExercisePaths(t *testing.T) {
	title := ExerciseTitlePath("ch01", "ex01_hash")
	if title != "ch01/exercises/ex01_hash_title.md" {
		t.Errorf("title path = %q", title)
	}
	desc := ExerciseDescPath("ch01", "ex01_hash")
	if desc != "ch01/exercises/ex01_hash_description.md" {
		t.Errorf("desc path = %q", desc)
	}
}

func TestResolver_ReadContent_English(t *testing.T) {
	contentFS := fstest.MapFS{
		"ch01/lesson01_01.md": &fstest.MapFile{Data: []byte("English text")},
	}
	r := NewResolver(contentFS, nil, "en")
	got, err := r.ReadContent("ch01/lesson01_01.md")
	if err != nil {
		t.Fatal(err)
	}
	if got != "English text" {
		t.Errorf("got %q, want %q", got, "English text")
	}
}

func TestResolver_ReadContent_LocaleFallback(t *testing.T) {
	contentFS := fstest.MapFS{
		"ch01/lesson01_01.md": &fstest.MapFile{Data: []byte("English text")},
	}
	localeFS := fstest.MapFS{} // Spanish has no translation.
	r := NewResolver(contentFS, localeFS, "es")
	got, err := r.ReadContent("ch01/lesson01_01.md")
	if err != nil {
		t.Fatal(err)
	}
	if got != "English text" {
		t.Errorf("got %q, want English fallback", got)
	}
}

func TestResolver_ReadContent_LocaleOverride(t *testing.T) {
	contentFS := fstest.MapFS{
		"ch01/lesson01_01.md": &fstest.MapFile{Data: []byte("English")},
	}
	localeFS := fstest.MapFS{
		"es/ch01/lesson01_01.md": &fstest.MapFile{Data: []byte("Spanish")},
	}
	r := NewResolver(contentFS, localeFS, "es")
	got, err := r.ReadContent("ch01/lesson01_01.md")
	if err != nil {
		t.Fatal(err)
	}
	if got != "Spanish" {
		t.Errorf("got %q, want locale override", got)
	}
}

func TestResolver_ReadContent_MissingFile(t *testing.T) {
	contentFS := fstest.MapFS{}
	r := NewResolver(contentFS, nil, "en")
	_, err := r.ReadContent("ch01/nonexistent.md")
	if err == nil {
		t.Error("expected error for missing file")
	}
}
