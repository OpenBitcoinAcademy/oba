// Package math provides a LaTeX subset parser and Gio-native renderer
// for mathematical formulas used in Bitcoin education content.
package math

// Greek maps LaTeX Greek letter commands to their Unicode codepoints.
var Greek = map[string]rune{
	`\alpha`:   'α',
	`\beta`:    'β',
	`\gamma`:   'γ',
	`\delta`:   'δ',
	`\epsilon`: 'ε',
	`\zeta`:    'ζ',
	`\eta`:     'η',
	`\theta`:   'θ',
	`\iota`:    'ι',
	`\kappa`:   'κ',
	`\lambda`:  'λ',
	`\mu`:      'μ',
	`\nu`:      'ν',
	`\xi`:      'ξ',
	`\pi`:      'π',
	`\rho`:     'ρ',
	`\sigma`:   'σ',
	`\tau`:     'τ',
	`\phi`:     'φ',
	`\chi`:     'χ',
	`\psi`:     'ψ',
	`\omega`:   'ω',
	// Uppercase variants used in Bitcoin math.
	`\Gamma`: 'Γ',
	`\Delta`: 'Δ',
	`\Sigma`: 'Σ',
	`\Pi`:    'Π',
	`\Omega`: 'Ω',
}

// Operators maps LaTeX operator commands to their Unicode codepoints.
var Operators = map[string]rune{
	`\cdot`:  '·',
	`\times`: '×',
	`\div`:   '÷',
	`\pm`:    '±',
	`\neq`:   '≠',
	`\equiv`: '≡',
	`\leq`:   '≤',
	`\geq`:   '≥',
	`\mod`:   0, // rendered as text "mod"
}

// Misc maps other LaTeX commands to their Unicode codepoints or rendering hints.
var Misc = map[string]rune{
	`\ldots`:  '…',
	`\infty`:  '∞',
	`\approx`: '≈',
	`\sum`:    '∑',
	`\prod`:   '∏',
	`\in`:     '∈',
	`\notin`:  '∉',
	`\forall`: '∀',
	`\exists`: '∃',
	`\to`:     '→',
}

// LookupSymbol returns the Unicode rune for a LaTeX command.
// Returns 0 and false if the command is not in any symbol table.
func LookupSymbol(cmd string) (rune, bool) {
	if r, ok := Greek[cmd]; ok {
		return r, true
	}
	if r, ok := Operators[cmd]; ok {
		return r, true
	}
	if r, ok := Misc[cmd]; ok {
		return r, true
	}
	return 0, false
}

// StructuralCommands maps LaTeX commands that affect layout rather than
// producing a single symbol (frac, sqrt, text, left, right).
var StructuralCommands = map[string]bool{
	`\frac`:  true,
	`\sqrt`:  true,
	`\text`:  true,
	`\left`:  true,
	`\right`: true,
}
