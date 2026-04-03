package math

import "strings"

// Unicode superscript digits.
var superDigits = map[rune]rune{
	'0': '\u2070', '1': '\u00B9', '2': '\u00B2', '3': '\u00B3',
	'4': '\u2074', '5': '\u2075', '6': '\u2076', '7': '\u2077',
	'8': '\u2078', '9': '\u2079',
}

// Unicode subscript digits.
var subDigits = map[rune]rune{
	'0': '\u2080', '1': '\u2081', '2': '\u2082', '3': '\u2083',
	'4': '\u2084', '5': '\u2085', '6': '\u2086', '7': '\u2087',
	'8': '\u2088', '9': '\u2089',
}

// Unicode subscript letters (limited set available in Unicode).
var subLetters = map[rune]rune{
	'a': '\u2090', 'e': '\u2091', 'h': '\u2095', 'i': '\u1D62',
	'j': '\u2C7C', 'k': '\u2096', 'l': '\u2097', 'm': '\u2098',
	'n': '\u2099', 'o': '\u2092', 'p': '\u209A', 'r': '\u1D63',
	's': '\u209B', 't': '\u209C', 'x': '\u2093',
}

// ToUnicode converts a LaTeX math expression to plain Unicode text.
// Returns the converted string and true if the expression is simple
// enough (digits, variables, superscripts, subscripts, symbols).
// Returns ("", false) for complex expressions (fractions, roots, etc.).
func ToUnicode(expr string) (string, bool) {
	if expr == "" {
		return "", false
	}
	node, err := Parse(expr)
	if err != nil {
		return "", false
	}
	return nodeToUnicode(node)
}

func nodeToUnicode(n *Node) (string, bool) {
	switch n.Type {
	case NodeText:
		return n.Value, true

	case NodeSymbol:
		// Parser already resolves symbols to Unicode characters.
		if n.Value == "" {
			return "", false
		}
		return n.Value, true

	case NodeSup:
		if len(n.Children) < 2 {
			return "", false
		}
		base, ok := nodeToUnicode(n.Children[0])
		if !ok {
			return "", false
		}
		sup, ok := toSuperscript(n.Children[1])
		if !ok {
			return "", false
		}
		return base + sup, true

	case NodeSub:
		if len(n.Children) < 2 {
			return "", false
		}
		base, ok := nodeToUnicode(n.Children[0])
		if !ok {
			return "", false
		}
		sub, ok := toSubscript(n.Children[1])
		if !ok {
			return "", false
		}
		return base + sub, true

	case NodeRow:
		var b strings.Builder
		for _, child := range n.Children {
			s, ok := nodeToUnicode(child)
			if !ok {
				return "", false
			}
			b.WriteString(s)
		}
		return b.String(), true

	case NodeGroup:
		if len(n.Children) == 1 {
			return nodeToUnicode(n.Children[0])
		}
		var b strings.Builder
		for _, child := range n.Children {
			s, ok := nodeToUnicode(child)
			if !ok {
				return "", false
			}
			b.WriteString(s)
		}
		return b.String(), true

	default:
		// NodeFrac, NodeSqrt, NodeDelimited, NodeSupSub, NodeTextCmd, NodeMod
		return "", false
	}
}

// toSuperscript converts a node's text content to Unicode superscript.
func toSuperscript(n *Node) (string, bool) {
	text, ok := flattenText(n)
	if !ok {
		return "", false
	}
	var b strings.Builder
	for _, r := range text {
		if s, ok := superDigits[r]; ok {
			b.WriteRune(s)
		} else {
			return "", false
		}
	}
	return b.String(), true
}

// toSubscript converts a node's text content to Unicode subscript.
func toSubscript(n *Node) (string, bool) {
	text, ok := flattenText(n)
	if !ok {
		return "", false
	}
	var b strings.Builder
	for _, r := range text {
		if s, ok := subDigits[r]; ok {
			b.WriteRune(s)
		} else if s, ok := subLetters[r]; ok {
			b.WriteRune(s)
		} else {
			return "", false
		}
	}
	return b.String(), true
}

// flattenText extracts plain text from a node tree.
// Only succeeds for NodeText, NodeGroup, and NodeRow containing only text.
func flattenText(n *Node) (string, bool) {
	switch n.Type {
	case NodeText:
		return n.Value, true
	case NodeGroup, NodeRow:
		var b strings.Builder
		for _, child := range n.Children {
			s, ok := flattenText(child)
			if !ok {
				return "", false
			}
			b.WriteString(s)
		}
		return b.String(), true
	default:
		return "", false
	}
}
