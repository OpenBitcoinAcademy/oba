package content

import "testing"

func TestTextSection_Validate(t *testing.T) {
	s := &TextSection{Content: "Hello"}
	if err := s.Validate(); err != nil {
		t.Errorf("valid text: %v", err)
	}
	if s.SectionType() != "text" {
		t.Errorf("type = %q, want %q", s.SectionType(), "text")
	}
}

func TestTextSection_ValidateEmpty(t *testing.T) {
	s := &TextSection{}
	if err := s.Validate(); err == nil {
		t.Error("empty text should fail validation")
	}
}

func TestCalloutSection_Validate(t *testing.T) {
	tests := []struct {
		style string
		ok    bool
	}{
		{"info", true},
		{"warning", true},
		{"tip", true},
		{"", false},
		{"error", false},
	}
	for _, tt := range tests {
		s := &CalloutSection{Content: "text", Style: tt.style}
		err := s.Validate()
		if tt.ok && err != nil {
			t.Errorf("style %q: unexpected error: %v", tt.style, err)
		}
		if !tt.ok && err == nil {
			t.Errorf("style %q: expected error", tt.style)
		}
	}
}

func TestCalloutSection_ValidateEmptyContent(t *testing.T) {
	s := &CalloutSection{Style: "info"}
	if err := s.Validate(); err == nil {
		t.Error("empty callout content should fail")
	}
}

func TestDiagramSection_Validate(t *testing.T) {
	s := &DiagramSection{DiagramID: "hash_flow"}
	if err := s.Validate(); err != nil {
		t.Errorf("valid diagram: %v", err)
	}
	s = &DiagramSection{}
	if err := s.Validate(); err == nil {
		t.Error("empty diagram_id should fail")
	}
}

func TestQuizSection_Validate(t *testing.T) {
	s := &QuizSection{
		Question: "What?",
		Options:  []string{"A", "B", "C"},
		Correct:  []int{1},
	}
	if err := s.Validate(); err != nil {
		t.Errorf("valid quiz: %v", err)
	}
}

func TestQuizSection_ValidateMultiAnswer(t *testing.T) {
	s := &QuizSection{
		Question: "Select all:",
		Options:  []string{"A", "B", "C", "D"},
		Correct:  []int{0, 2, 3},
	}
	if err := s.Validate(); err != nil {
		t.Errorf("valid multi-answer quiz: %v", err)
	}
}

func TestQuizSection_ValidateErrors(t *testing.T) {
	tests := []struct {
		name string
		q    QuizSection
	}{
		{"empty question", QuizSection{Options: []string{"A", "B"}, Correct: []int{0}}},
		{"one option", QuizSection{Question: "Q", Options: []string{"A"}, Correct: []int{0}}},
		{"no correct", QuizSection{Question: "Q", Options: []string{"A", "B"}}},
		{"index out of range", QuizSection{Question: "Q", Options: []string{"A", "B"}, Correct: []int{5}}},
		{"negative index", QuizSection{Question: "Q", Options: []string{"A", "B"}, Correct: []int{-1}}},
	}
	for _, tt := range tests {
		if err := tt.q.Validate(); err == nil {
			t.Errorf("%s: expected error", tt.name)
		}
	}
}

func TestInteractiveSection_Validate(t *testing.T) {
	s := &InteractiveSection{ExerciseID: "ex01_hash"}
	if err := s.Validate(); err != nil {
		t.Errorf("valid interactive: %v", err)
	}
	s = &InteractiveSection{}
	if err := s.Validate(); err == nil {
		t.Error("empty exercise_id should fail")
	}
}

func TestFormulaSection_Validate(t *testing.T) {
	s := &FormulaSection{Formula: "y^2 = x^3 + 7"}
	if err := s.Validate(); err != nil {
		t.Errorf("valid formula: %v", err)
	}
	s = &FormulaSection{}
	if err := s.Validate(); err == nil {
		t.Error("empty formula should fail")
	}
}

func TestCodeSection_Validate(t *testing.T) {
	s := &CodeSection{Content: "print('hello')", Language: "python"}
	if err := s.Validate(); err != nil {
		t.Errorf("valid code: %v", err)
	}
	s = &CodeSection{Language: "go"}
	if err := s.Validate(); err == nil {
		t.Error("empty code content should fail")
	}
}
