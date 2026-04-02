package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ScriptExecution shows the P2PKH script flow with a stack visualization.
// Left: opcode sequence. Right: stack state after OP_DUP.
type ScriptExecution struct{}

func (d *ScriptExecution) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(220))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	opcodeW := pct(w, 50)
	stackW := pct(w, 35)
	gap := w - 2*pad - opcodeW - stackW
	if gap < 8 {
		gap = 8
	}

	bh := gtx.Dp(unit.Dp(24))
	spacing := 4

	// Left column: opcode sequence.
	ox := pad
	oy := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.script_opcodes"), image.Pt(ox, oy), th.Color.TextMuted)
	oy += gtx.Dp(unit.Dp(18))

	opcodes := []struct {
		label string
		isOp  bool
	}{
		{i18n.T("diagram.script_sig_label"), false},
		{i18n.T("diagram.script_pubkey_label"), false},
		{"OP_DUP", true},
		{"OP_HASH160", true},
		{i18n.T("diagram.script_hash_label"), false},
		{"OP_EQUALVERIFY", true},
		{"OP_CHECKSIG", true},
	}

	for _, op := range opcodes {
		if op.isOp {
			processBox(gtx, th, op.label, image.Pt(ox, oy), opcodeW, bh)
		} else {
			box(gtx, th, op.label, image.Pt(ox, oy), opcodeW, bh, th.Color.InfoBg)
		}
		oy += bh + spacing
	}

	// Right column: stack state (after OP_DUP as example).
	sx := pad + opcodeW + gap
	sy := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.script_stack_label")+" (after OP_DUP)", image.Pt(sx, sy), th.Color.TextMuted)
	sy += gtx.Dp(unit.Dp(18))

	stackItems := []string{
		i18n.T("diagram.script_pubkey_label"),
		i18n.T("diagram.script_pubkey_label"),
		i18n.T("diagram.script_sig_label"),
	}

	// Draw stack bottom-up visually (top of stack at top).
	for _, item := range stackItems {
		box(gtx, th, item, image.Pt(sx, sy), stackW, bh, th.Color.TipBg)
		sy += bh + spacing
	}

	// Arrow from OP_DUP to stack.
	dupY := gtx.Dp(unit.Dp(8)) + gtx.Dp(unit.Dp(18)) + 2*(bh+spacing) + bh/2
	arrow(gtx, th, image.Pt(ox+opcodeW, dupY), image.Pt(sx, dupY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
