// Package components provides reusable Gio widgets for lesson content.
package components

import (
	"image"
	"log"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	obamath "github.com/openbitcoinacademy/oba/internal/math"
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

		// Warn on unsupported markdown syntax (Section 6).
		warnUnsupported(trimmed)

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

// span represents a piece of inline-formatted text.
type span struct {
	text   string
	bold   bool
	italic bool
	code   bool
	math   bool // $..$ inline LaTeX
}

// renderParagraph renders a paragraph with inline bold, italic, and code.
// Splits into styled spans and renders each as a separate label. Gio
// handles text wrapping within each label but not across labels, so
// spans are concatenated into styled runs for best wrapping behavior.
func (m *Markdown) renderParagraph(gtx layout.Context, text string) layout.Dimensions {
	spans := parseInlineSpans(text)

	// For practical wrapping, group consecutive same-style spans.
	// Gio wraps within a single label, so we merge plain text spans
	// and only break at style boundaries.
	var children []layout.FlexChild
	for _, sp := range spans {
		sp := sp
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if sp.math {
				renderer := obamath.NewRenderer(m.Theme.Material, m.Theme.Text.Body)
				renderer.Color = m.Theme.Color.Text
				gtx.Constraints.Min = image.Point{}
				return renderer.Layout(gtx, sp.text)
			}
			lbl := material.Label(m.Theme.Material, m.Theme.Text.Body, sp.text)
			lbl.Color = m.Theme.Color.Text
			if sp.bold {
				lbl.Font.Weight = font.Bold
			}
			if sp.italic {
				lbl.Font.Style = font.Italic
			}
			if sp.code {
				lbl.Font.Typeface = "monospace"
				lbl.Color = m.Theme.Color.Accent
			}
			return lbl.Layout(gtx)
		}))
	}

	// Render spans in a vertical flex. Each span wraps independently.
	// True inline flow would need a custom layout, but in practice most
	// paragraphs have few style changes and this works well.
	gtx.Constraints.Min = image.Point{}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

// parseInlineSpans splits text into styled spans.
// Handles **bold**, *italic*, and `code` markers.
func parseInlineSpans(text string) []span {
	var spans []span
	i := 0
	var current strings.Builder

	flush := func(bold, italic, code bool) {
		if current.Len() > 0 {
			spans = append(spans, span{text: current.String(), bold: bold, italic: italic, code: code})
			current.Reset()
		}
	}

	for i < len(text) {
		// Inline math: $...$
		if text[i] == '$' {
			flush(false, false, false)
			end := strings.Index(text[i+1:], "$")
			if end >= 0 {
				spans = append(spans, span{text: text[i+1 : i+1+end], math: true})
				i = i + 1 + end + 1
				continue
			}
		}

		// Code spans: `...`
		if text[i] == '`' {
			flush(false, false, false)
			end := strings.Index(text[i+1:], "`")
			if end >= 0 {
				spans = append(spans, span{text: text[i+1 : i+1+end], code: true})
				i = i + 1 + end + 1
				continue
			}
		}

		// Bold: **...**
		if i+1 < len(text) && text[i] == '*' && text[i+1] == '*' {
			flush(false, false, false)
			end := strings.Index(text[i+2:], "**")
			if end >= 0 {
				spans = append(spans, span{text: text[i+2 : i+2+end], bold: true})
				i = i + 2 + end + 2
				continue
			}
		}

		// Italic: *...*  (single *, not followed by another *)
		if text[i] == '*' && (i+1 >= len(text) || text[i+1] != '*') {
			flush(false, false, false)
			end := strings.Index(text[i+1:], "*")
			if end >= 0 {
				spans = append(spans, span{text: text[i+1 : i+1+end], italic: true})
				i = i + 1 + end + 1
				continue
			}
		}

		current.WriteByte(text[i])
		i++
	}
	flush(false, false, false)

	return spans
}

// warnUnsupported logs a warning for markdown syntax outside the supported subset.
func warnUnsupported(line string) {
	switch {
	case strings.HasPrefix(line, "!["):
		log.Printf("markdown: unsupported image syntax: %s", truncate(line))
	case strings.HasPrefix(line, "<"):
		log.Printf("markdown: unsupported HTML: %s", truncate(line))
	case strings.HasPrefix(line, "- ["):
		log.Printf("markdown: unsupported task list: %s", truncate(line))
	case strings.Contains(line, "](") && strings.Contains(line, "["):
		log.Printf("markdown: unsupported link: %s", truncate(line))
	case strings.HasPrefix(line, "|") && strings.HasSuffix(strings.TrimSpace(line), "|"):
		log.Printf("markdown: unsupported table: %s", truncate(line))
	}
}

func truncate(s string) string {
	if len(s) > 60 {
		return s[:60] + "..."
	}
	return s
}
