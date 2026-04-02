package content

import (
	"fmt"
	"io/fs"

	"github.com/BurntSushi/toml"
)

// ExerciseConfig holds the parsed configuration for an exercise.
// The Type field determines which widget renders it. The Config
// map holds type-specific settings from the TOML file.
type ExerciseConfig struct {
	ID     string
	Type   string
	Config map[string]interface{}
}

// rawExercise mirrors the TOML structure.
type rawExercise struct {
	Exercise struct {
		ID     string                 `toml:"id"`
		Type   string                 `toml:"type"`
		Config map[string]interface{} `toml:"config"`
	} `toml:"exercise"`
}

// LoadExercise parses an exercise TOML file.
func LoadExercise(data []byte, path string) (*ExerciseConfig, error) {
	var raw rawExercise
	if err := toml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if raw.Exercise.ID == "" {
		return nil, fmt.Errorf("%s: exercise missing id", path)
	}
	if raw.Exercise.Type == "" {
		return nil, fmt.Errorf("%s: exercise missing type", path)
	}
	return &ExerciseConfig{
		ID:     raw.Exercise.ID,
		Type:   raw.Exercise.Type,
		Config: raw.Exercise.Config,
	}, nil
}

// LoadExerciseFromFS reads and parses an exercise TOML from a filesystem.
func LoadExerciseFromFS(fsys fs.FS, path string) (*ExerciseConfig, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return LoadExercise(data, path)
}

// ConfigString returns a string config value, or the default if missing.
func (ec *ExerciseConfig) ConfigString(key, def string) string {
	if v, ok := ec.Config[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return def
}

// ConfigBool returns a bool config value, or the default if missing.
func (ec *ExerciseConfig) ConfigBool(key string, def bool) bool {
	if v, ok := ec.Config[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return def
}

// ConfigInt returns an int64 config value, or the default if missing.
func (ec *ExerciseConfig) ConfigInt(key string, def int64) int64 {
	if v, ok := ec.Config[key]; ok {
		if i, ok := v.(int64); ok {
			return i
		}
	}
	return def
}

// LoadChapterExercises scans a chapter's exercises directory for .toml files
// and returns a map of exercise ID to config.
func LoadChapterExercises(fsys fs.FS, chapterID string) (map[string]*ExerciseConfig, error) {
	dir := chapterID + "/exercises"
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		// No exercises directory is valid (chapter might not have exercises).
		return nil, nil
	}

	exercises := make(map[string]*ExerciseConfig)
	for _, e := range entries {
		if e.IsDir() || !hasTomlExt(e.Name()) {
			continue
		}
		path := dir + "/" + e.Name()
		ec, err := LoadExerciseFromFS(fsys, path)
		if err != nil {
			return nil, err
		}
		exercises[ec.ID] = ec
	}
	return exercises, nil
}

func hasTomlExt(name string) bool {
	return len(name) > 5 && name[len(name)-5:] == ".toml"
}

// ConfigStrings returns a string slice config value, or nil if missing.
func (ec *ExerciseConfig) ConfigStrings(key string) []string {
	if v, ok := ec.Config[key]; ok {
		if arr, ok := v.([]interface{}); ok {
			var result []string
			for _, item := range arr {
				if s, ok := item.(string); ok {
					result = append(result, s)
				}
			}
			return result
		}
	}
	return nil
}
