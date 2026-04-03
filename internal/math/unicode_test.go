package math

import "testing"

func TestToUnicode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		ok    bool
	}{
		{"superscript digits", "2^{256}", "2²⁵⁶", true},
		{"superscript short", "2^{32}", "2³²", true},
		{"single variable", "i", "i", true},
		{"single uppercase", "S", "S", true},
		{"plain number", "256", "256", true},
		{"subscript digit", "x_0", "x₀", true},
		{"subscript letter", "x_i", "xᵢ", true},
		{"greek with subscript", `\lambda_i`, "λᵢ", true},
		{"greek letter", `\alpha`, "α", true},
		{"operator", `\times`, "×", true},
		{"frac rejected", `\frac{1}{2}`, "", false},
		{"sqrt rejected", `\sqrt{2}`, "", false},
		{"complex formula rejected", `\sum_{i=0}^{10} i`, "", false},
		{"empty string", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ToUnicode(tt.input)
			if ok != tt.ok {
				t.Errorf("ToUnicode(%q) ok = %v, want %v", tt.input, ok, tt.ok)
			}
			if got != tt.want {
				t.Errorf("ToUnicode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
