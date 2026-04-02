package content

import (
	"testing"
	"testing/fstest"
)

var hashToml = `
[exercise]
id = "ex01_hash"
type = "hash_explorer"

[exercise.config]
algorithm = "sha256"
preloaded_examples = ["bitcoin", "Bitcoin", "bitcoin!"]
output_formats = ["hex", "binary", "base64"]
debounce_ms = 100
`

func TestLoadExercise_Basic(t *testing.T) {
	ec, err := LoadExercise([]byte(hashToml), "test.toml")
	if err != nil {
		t.Fatal(err)
	}
	if ec.ID != "ex01_hash" {
		t.Errorf("ID = %q", ec.ID)
	}
	if ec.Type != "hash_explorer" {
		t.Errorf("Type = %q", ec.Type)
	}
	if ec.ConfigString("algorithm", "") != "sha256" {
		t.Errorf("algorithm = %q", ec.ConfigString("algorithm", ""))
	}
}

func TestLoadExercise_ConfigAccessors(t *testing.T) {
	ec, _ := LoadExercise([]byte(hashToml), "test.toml")

	examples := ec.ConfigStrings("preloaded_examples")
	if len(examples) != 3 {
		t.Fatalf("examples = %d, want 3", len(examples))
	}
	if examples[0] != "bitcoin" {
		t.Errorf("examples[0] = %q", examples[0])
	}

	debounce := ec.ConfigInt("debounce_ms", 0)
	if debounce != 100 {
		t.Errorf("debounce_ms = %d", debounce)
	}
}

func TestLoadExercise_BoolConfig(t *testing.T) {
	toml := `
[exercise]
id = "ex02_keys"
type = "key_generator"

[exercise.config]
show_binary_option = true
show_byte_count = true
`
	ec, err := LoadExercise([]byte(toml), "test.toml")
	if err != nil {
		t.Fatal(err)
	}
	if !ec.ConfigBool("show_binary_option", false) {
		t.Error("show_binary_option should be true")
	}
	if ec.ConfigBool("nonexistent", false) {
		t.Error("nonexistent should default to false")
	}
}

func TestLoadExercise_MissingID(t *testing.T) {
	bad := `[exercise]
type = "hash_explorer"
`
	_, err := LoadExercise([]byte(bad), "test.toml")
	if err == nil {
		t.Error("expected error for missing ID")
	}
}

func TestLoadExercise_MissingType(t *testing.T) {
	bad := `[exercise]
id = "ex01"
`
	_, err := LoadExercise([]byte(bad), "test.toml")
	if err == nil {
		t.Error("expected error for missing type")
	}
}

func TestLoadChapterExercises(t *testing.T) {
	fsys := fstest.MapFS{
		"ch01/exercises/ex01_hash.toml":     {Data: []byte(hashToml)},
		"ch01/exercises/ex01_hash_title.md": {Data: []byte("Hash Explorer\n")},
	}

	exercises, err := LoadChapterExercises(fsys, "ch01")
	if err != nil {
		t.Fatal(err)
	}
	if len(exercises) != 1 {
		t.Fatalf("exercises = %d, want 1", len(exercises))
	}
	if _, ok := exercises["ex01_hash"]; !ok {
		t.Error("missing ex01_hash")
	}
}

func TestLoadChapterExercises_NoDir(t *testing.T) {
	fsys := fstest.MapFS{}
	exercises, err := LoadChapterExercises(fsys, "ch99")
	if err != nil {
		t.Fatal(err)
	}
	if exercises != nil {
		t.Error("should return nil for missing dir")
	}
}

func TestConfigDefaults(t *testing.T) {
	ec := &ExerciseConfig{Config: map[string]interface{}{}}
	if ec.ConfigString("x", "default") != "default" {
		t.Error("string default")
	}
	if ec.ConfigInt("x", 42) != 42 {
		t.Error("int default")
	}
	if ec.ConfigBool("x", true) != true {
		t.Error("bool default")
	}
	if ec.ConfigStrings("x") != nil {
		t.Error("strings default")
	}
}
