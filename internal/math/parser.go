package math

import (
	"fmt"
	"strings"
	"unicode"
)

// NodeType identifies the kind of AST node.
type NodeType int

const (
	NodeRow       NodeType = iota // Horizontal sequence of nodes
	NodeText                      // Literal text characters
	NodeSymbol                    // Greek letter or operator from symbol table
	NodeSup                       // Superscript: Base^{Sup}
	NodeSub                       // Subscript: Base_{Sub}
	NodeSupSub                    // Both: Base^{Sup}_{Sub} or Base_{Sub}^{Sup}
	NodeFrac                      // \frac{Num}{Den}
	NodeSqrt                      // \sqrt{Content}
	NodeGroup                     // {Content} braced group
	NodeDelimited                 // \left( ... \right)
	NodeTextCmd                   // \text{...} literal text in math
	NodeMod                       // \mod rendered as "mod" with spacing
)

// Node is a single element in the math AST.
type Node struct {
	Type     NodeType
	Value    string  // For NodeText, NodeSymbol: the character(s)
	Children []*Node // For NodeRow, NodeGroup, NodeDelimited
	// Script nodes: Children[0]=base, Children[1]=sup/sub, Children[2]=second script
	// Frac nodes: Children[0]=numerator, Children[1]=denominator
	// Sqrt nodes: Children[0]=content
	Left  string // For NodeDelimited: opening delimiter
	Right string // For NodeDelimited: closing delimiter
}

// Parse converts a LaTeX subset string into a math AST.
// Returns an error for unsupported commands.
func Parse(input string) (*Node, error) {
	p := &parser{input: input}
	nodes, err := p.parseSequence(false)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 1 {
		return nodes[0], nil
	}
	return &Node{Type: NodeRow, Children: nodes}, nil
}

type parser struct {
	input string
	pos   int
}

func (p *parser) peek() (byte, bool) {
	if p.pos >= len(p.input) {
		return 0, false
	}
	return p.input[p.pos], true
}

func (p *parser) advance() byte {
	b := p.input[p.pos]
	p.pos++
	return b
}

func (p *parser) skipSpaces() {
	for p.pos < len(p.input) && p.input[p.pos] == ' ' {
		p.pos++
	}
}

// parseSequence reads nodes until EOF or a closing brace (if inGroup).
func (p *parser) parseSequence(inGroup bool) ([]*Node, error) {
	var nodes []*Node

	for {
		p.skipSpaces()
		ch, ok := p.peek()
		if !ok {
			break
		}

		if ch == '}' {
			if inGroup {
				break
			}
			return nil, fmt.Errorf("unexpected '}' at position %d", p.pos)
		}

		node, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		if node == nil {
			break
		}

		// Check for super/subscripts after any atom.
		node, err = p.parseScripts(node)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

// parseAtom reads a single element: character, command, or braced group.
func (p *parser) parseAtom() (*Node, error) {
	ch, ok := p.peek()
	if !ok {
		return nil, nil
	}

	switch {
	case ch == '\\':
		return p.parseCommand()
	case ch == '{':
		return p.parseGroup()
	case ch == '^' || ch == '_':
		// Bare script without a preceding base: treat as empty base.
		return &Node{Type: NodeText, Value: ""}, nil
	default:
		return p.parseChar()
	}
}

// parseChar reads a single non-command character.
func (p *parser) parseChar() (*Node, error) {
	b := p.advance()
	return &Node{Type: NodeText, Value: string(b)}, nil
}

// parseGroup reads { ... } and returns a group node.
func (p *parser) parseGroup() (*Node, error) {
	p.advance() // consume '{'
	nodes, err := p.parseSequence(true)
	if err != nil {
		return nil, err
	}
	ch, ok := p.peek()
	if !ok || ch != '}' {
		return nil, fmt.Errorf("unclosed group starting before position %d", p.pos)
	}
	p.advance() // consume '}'

	if len(nodes) == 1 {
		return nodes[0], nil
	}
	return &Node{Type: NodeGroup, Children: nodes}, nil
}

// parseCommand reads a backslash command.
func (p *parser) parseCommand() (*Node, error) {
	start := p.pos
	p.advance() // consume '\'

	// Read command name (letters only).
	cmdStart := p.pos
	for p.pos < len(p.input) && isLetter(p.input[p.pos]) {
		p.pos++
	}

	if p.pos == cmdStart {
		return nil, fmt.Errorf("empty command at position %d", start)
	}

	cmd := `\` + p.input[cmdStart:p.pos]

	// Structural commands.
	if StructuralCommands[cmd] {
		return p.parseStructural(cmd)
	}

	// Symbol lookup.
	if r, ok := LookupSymbol(cmd); ok {
		if cmd == `\mod` {
			return &Node{Type: NodeMod, Value: "mod"}, nil
		}
		return &Node{Type: NodeSymbol, Value: string(r)}, nil
	}

	return nil, fmt.Errorf("unsupported command %q at position %d", cmd, start)
}

// parseStructural handles \frac, \sqrt, \text, \left, \right.
func (p *parser) parseStructural(cmd string) (*Node, error) {
	switch cmd {
	case `\frac`:
		return p.parseFrac()
	case `\sqrt`:
		return p.parseSqrt()
	case `\text`:
		return p.parseTextCmd()
	case `\left`:
		return p.parseDelimited()
	case `\right`:
		// \right without matching \left is an error.
		return nil, fmt.Errorf(`\right without matching \left at position %d`, p.pos)
	default:
		return nil, fmt.Errorf("unknown structural command %q", cmd)
	}
}

func (p *parser) parseFrac() (*Node, error) {
	p.skipSpaces()
	num, err := p.parseBracedArg()
	if err != nil {
		return nil, fmt.Errorf("\\frac numerator: %w", err)
	}
	p.skipSpaces()
	den, err := p.parseBracedArg()
	if err != nil {
		return nil, fmt.Errorf("\\frac denominator: %w", err)
	}
	return &Node{Type: NodeFrac, Children: []*Node{num, den}}, nil
}

func (p *parser) parseSqrt() (*Node, error) {
	p.skipSpaces()
	content, err := p.parseBracedArg()
	if err != nil {
		return nil, fmt.Errorf("\\sqrt argument: %w", err)
	}
	return &Node{Type: NodeSqrt, Children: []*Node{content}}, nil
}

func (p *parser) parseTextCmd() (*Node, error) {
	p.skipSpaces()
	ch, ok := p.peek()
	if !ok || ch != '{' {
		return nil, fmt.Errorf("\\text requires braced argument at position %d", p.pos)
	}
	p.advance() // consume '{'

	var text strings.Builder
	for {
		ch, ok := p.peek()
		if !ok {
			return nil, fmt.Errorf("unclosed \\text at position %d", p.pos)
		}
		if ch == '}' {
			p.advance()
			break
		}
		text.WriteByte(p.advance())
	}
	return &Node{Type: NodeTextCmd, Value: text.String()}, nil
}

func (p *parser) parseDelimited() (*Node, error) {
	// Read opening delimiter.
	p.skipSpaces()
	left, err := p.readDelimiter()
	if err != nil {
		return nil, fmt.Errorf("\\left delimiter: %w", err)
	}

	// Parse content until \right.
	var nodes []*Node
	for {
		p.skipSpaces()
		ch, ok := p.peek()
		if !ok {
			return nil, fmt.Errorf("unclosed \\left at position %d", p.pos)
		}
		if ch == '\\' && p.looksLikeRight() {
			break
		}

		node, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		if node == nil {
			break
		}
		node, err = p.parseScripts(node)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	// Consume \right.
	p.advance() // '\'
	cmdStart := p.pos
	for p.pos < len(p.input) && isLetter(p.input[p.pos]) {
		p.pos++
	}
	cmd := p.input[cmdStart:p.pos]
	if cmd != "right" {
		return nil, fmt.Errorf("expected \\right, got \\%s", cmd)
	}

	p.skipSpaces()
	right, err := p.readDelimiter()
	if err != nil {
		return nil, fmt.Errorf("\\right delimiter: %w", err)
	}

	content := &Node{Type: NodeRow, Children: nodes}
	return &Node{
		Type:     NodeDelimited,
		Children: []*Node{content},
		Left:     left,
		Right:    right,
	}, nil
}

func (p *parser) looksLikeRight() bool {
	if p.pos+6 > len(p.input) {
		return false
	}
	return p.input[p.pos:p.pos+6] == `\right`
}

func (p *parser) readDelimiter() (string, error) {
	ch, ok := p.peek()
	if !ok {
		return "", fmt.Errorf("expected delimiter at position %d", p.pos)
	}
	switch ch {
	case '(', ')', '[', ']', '|', '.':
		p.advance()
		return string(ch), nil
	default:
		return "", fmt.Errorf("unsupported delimiter %q at position %d", string(ch), p.pos)
	}
}

// parseBracedArg reads a { ... } argument and returns its content as a node.
func (p *parser) parseBracedArg() (*Node, error) {
	ch, ok := p.peek()
	if !ok || ch != '{' {
		return nil, fmt.Errorf("expected '{' at position %d", p.pos)
	}
	return p.parseGroup()
}

// parseScripts checks for ^ and _ after a base node and attaches them.
func (p *parser) parseScripts(base *Node) (*Node, error) {
	var sup, sub *Node

	for {
		p.skipSpaces()
		ch, ok := p.peek()
		if !ok {
			break
		}

		if ch == '^' && sup == nil {
			p.advance()
			arg, err := p.parseScriptArg()
			if err != nil {
				return nil, fmt.Errorf("superscript: %w", err)
			}
			sup = arg
		} else if ch == '_' && sub == nil {
			p.advance()
			arg, err := p.parseScriptArg()
			if err != nil {
				return nil, fmt.Errorf("subscript: %w", err)
			}
			sub = arg
		} else {
			break
		}
	}

	if sup != nil && sub != nil {
		return &Node{Type: NodeSupSub, Children: []*Node{base, sup, sub}}, nil
	}
	if sup != nil {
		return &Node{Type: NodeSup, Children: []*Node{base, sup}}, nil
	}
	if sub != nil {
		return &Node{Type: NodeSub, Children: []*Node{base, sub}}, nil
	}
	return base, nil
}

// parseScriptArg reads either a braced group or a single character/command
// as a script argument.
func (p *parser) parseScriptArg() (*Node, error) {
	p.skipSpaces()
	ch, ok := p.peek()
	if !ok {
		return nil, fmt.Errorf("expected script argument at position %d", p.pos)
	}
	if ch == '{' {
		return p.parseGroup()
	}
	// Single token.
	return p.parseAtom()
}

func isLetter(b byte) bool {
	return unicode.IsLetter(rune(b))
}
