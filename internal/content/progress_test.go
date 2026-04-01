package content

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestNewProgress(t *testing.T) {
	p := NewProgress()
	if p.Version != 1 {
		t.Errorf("version = %d", p.Version)
	}
	if p.Locale != "en" {
		t.Errorf("locale = %q", p.Locale)
	}
	if len(p.Completed) != 0 {
		t.Error("completed should be empty")
	}
}

func TestProgress_MarkAndCheck(t *testing.T) {
	p := NewProgress()
	if p.IsLessonComplete("ch01.lesson01") {
		t.Error("should not be complete yet")
	}
	p.MarkLessonComplete("ch01.lesson01")
	if !p.IsLessonComplete("ch01.lesson01") {
		t.Error("should be complete")
	}
	if p.CompletedCount() != 1 {
		t.Errorf("count = %d", p.CompletedCount())
	}
}

func TestProgress_ExerciseMarkAndCheck(t *testing.T) {
	p := NewProgress()
	p.MarkExerciseComplete("ex01_hash")
	if !p.IsExerciseComplete("ex01_hash") {
		t.Error("exercise should be complete")
	}
	if p.IsExerciseComplete("ex02_keys") {
		t.Error("other exercise should not be complete")
	}
}

func TestProgress_Reset(t *testing.T) {
	p := NewProgress()
	p.MarkLessonComplete("ch01.lesson01")
	p.MarkExerciseComplete("ex01_hash")
	p.Reset()
	if p.CompletedCount() != 0 {
		t.Error("reset should clear lessons")
	}
	if p.IsExerciseComplete("ex01_hash") {
		t.Error("reset should clear exercises")
	}
}

func TestProgress_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")

	p := NewProgress()
	p.Locale = "es"
	p.MarkLessonComplete("ch01.lesson01")
	p.MarkExerciseComplete("ex01_hash")

	if err := SaveProgress(path, p); err != nil {
		t.Fatal(err)
	}

	loaded := LoadProgress(path)
	if loaded.Locale != "es" {
		t.Errorf("locale = %q", loaded.Locale)
	}
	if !loaded.IsLessonComplete("ch01.lesson01") {
		t.Error("lesson should be complete after load")
	}
	if !loaded.IsExerciseComplete("ex01_hash") {
		t.Error("exercise should be complete after load")
	}
}

func TestLoadProgress_MissingFile(t *testing.T) {
	p := LoadProgress("/nonexistent/path/progress.json")
	if p.Version != 1 {
		t.Error("missing file should return fresh progress")
	}
}

func TestLoadProgress_CorruptedFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")
	os.WriteFile(path, []byte("not json{{{"), 0o644)

	p := LoadProgress(path)
	if p.Version != 1 {
		t.Error("corrupted file should return fresh progress")
	}
}

func TestLoadProgress_FutureVersion(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")
	os.WriteFile(path, []byte(`{"version":999}`), 0o644)

	p := LoadProgress(path)
	if p.Version != 1 {
		t.Error("future version should reset to current")
	}
}

func TestLoadProgressFS(t *testing.T) {
	fsys := fstest.MapFS{
		"progress.json": {Data: []byte(`{"version":1,"locale":"pt","completed":{"ch01.lesson01":"2024-01-01T00:00:00Z"},"exercises":{}}`)},
	}
	p := LoadProgressFS(fsys, "progress.json")
	if p.Locale != "pt" {
		t.Errorf("locale = %q", p.Locale)
	}
	if !p.IsLessonComplete("ch01.lesson01") {
		t.Error("should be complete")
	}
}

func TestLoadProgressFS_Missing(t *testing.T) {
	fsys := fstest.MapFS{}
	p := LoadProgressFS(fsys, "progress.json")
	if p.Version != 1 {
		t.Error("missing should return fresh")
	}
}
