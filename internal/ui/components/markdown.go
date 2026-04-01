// Package components provides reusable Gio widgets for lesson content.
package components

import (
	"image"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Markdown renders a subset of Markdown as Gio widgets.
// Supported: paragraphs, **bold**, *italic*, `code`, ## headings, ``` code blocks.
type Markdown struct {
	Content string
	Theme   *theme.Theme
}

// Layout renders the markdown content.
func (m *Markdown) Layout(gtx layout.Context) layout.Dimensions {
	blocks := parseBlocks(m.Content)
	var children []layout.FlexChild

	for _, block := range blocks {
		b := block // capture
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.renderBlock(gtx, b)
		}))
		children = append(children, layout.Rigid(layout.Spacer{Height: m.Theme.Space.Small}.Layout))
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

type block struct {
	kind    string // "para", "h2", "h3", "code"
	content string
}

func parseBlocks(content string) []block {
	var blocks []block
	lines := strings.Split(content, "\n")
	var current []string
	inCode := false

	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			if inCode {
				blocks = append(blocks, block{kind: "code", content: strings.Join(current, "\n")})
				current = nil
				inCode = false
			} else {
				if len(current) > 0 {
					blocks = append(blocks, block{kind: "para", content: strings.Join(current, " ")})
					current = nil
				}
				inCode = true
			}
			continue
		}

		if inCode {
			current = append(current, line)
			continue
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if len(current) > 0 {
				blocks = append(blocks, block{kind: "para", content: strings.Join(current, " ")})
				current = nil
			}
			continue
		}

		if strings.HasPrefix(trimmed, "### ") {
			if len(current) > 0 {
				blocks = append(blocks, block{kind: "para", content: strings.Join(current, " ")})
				current = nil
			}
			blocks = append(blocks, block{kind: "h3", content: strings.TrimPrefix(trimmed, "### ")})
			continue
		}
		if strings.HasPrefix(trimmed, "## ") {
			if len(current) > 0 {
				blocks = append(blocks, block{kind: "para", content: strings.Join(current, " ")})
				current = nil
			}
			blocks = append(blocks, block{kind: "h2", content: strings.TrimPrefix(trimmed, "## ")})
			continue
		}

		current = append(current, trimmed)
	}

	if len(current) > 0 {
		if inCode {
			blocks = append(blocks, block{kind: "code", content: strings.Join(current, "\n")})
		} else {
			blocks = append(blocks, block{kind: "para", content: strings.Join(current, " ")})
		}
	}

	return blocks
}

func (m *Markdown) renderBlock(gtx layout.Context, b block) layout.Dimensions {
	switch b.kind {
	case "h2":
		lbl := material.Label(m.Theme.Material, m.Theme.Text.H2, b.content)
		lbl.Color = m.Theme.Color.Text
		lbl.Font.Weight = font.Bold
		return lbl.Layout(gtx)
	case "h3":
		lbl := material.Label(m.Theme.Material, m.Theme.Text.H3, b.content)
		lbl.Color = m.Theme.Color.Text
		lbl.Font.Weight = font.Bold
		return lbl.Layout(gtx)
	case "code":
		lbl := material.Label(m.Theme.Material, m.Theme.Text.Code, b.content)
		lbl.Color = m.Theme.Color.Text
		lbl.Font.Typeface = "monospace"
		return layout.Inset{
			Top: unit.Dp(4), Bottom: unit.Dp(4),
			Left: unit.Dp(8), Right: unit.Dp(8),
		}.Layout(gtx, lbl.Layout)
	default:
		return m.renderParagraph(gtx, b.content)
	}
}

// renderParagraph renders a paragraph with inline formatting.
// For v1, inline bold/italic/code are stripped to plain text.
// Full inline formatting will be added when Gio's rich text API matures.
func (m *Markdown) renderParagraph(gtx layout.Context, text string) layout.Dimensions {
	// Strip inline markdown for now (render as plain text).
	plain := stripInlineMarkdown(text)
	lbl := material.Label(m.Theme.Material, m.Theme.Text.Body, plain)
	lbl.Color = m.Theme.Color.Text
	gtx.Constraints.Min = image.Point{}
	return lbl.Layout(gtx)
}

func stripInlineMarkdown(s string) string {
	// Remove **bold** markers.
	s = strings.ReplaceAll(s, "**", "")
	// Remove *italic* markers (but not inside words like "don't").
	// Simple approach: remove standalone * markers.
	result := strings.Builder{}
	prev := byte(0)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' && (i+1 >= len(s) || s[i+1] == ' ' || prev == ' ' || prev == 0) {
			continue
		}
		result.WriteByte(s[i])
		prev = s[i]
	}
	// Remove `code` backtick markers.
	return strings.ReplaceAll(result.String(), "`", "")
}
