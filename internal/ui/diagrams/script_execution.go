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
	h := gtx.Dp(unit.Dp(240))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	opcodeW := pct(w, 48)
	stackW := pct(w, 36)
	gap := w - 2*pad - opcodeW - stackW
	if gap < 10 {
		gap = 10
	}

	bh := gtx.Dp(unit.Dp(24))
	spacing := 5

	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Left column: opcode sequence.
	ox := pad
	oy := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.script_opcodes"), image.Pt(ox, oy), th.Color.TextMuted)
	oy += gtx.Dp(unit.Dp(24))

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

	dupY := 0 // track OP_DUP's Y for the arrow
	for idx, op := range opcodes {
		if op.isOp {
			shadowBox(gtx, th, op.label, image.Pt(ox, oy), opcodeW, bh, th.Color.Primary)
		} else {
			box(gtx, th, op.label, image.Pt(ox, oy), opcodeW, bh, th.Color.InfoBg)
		}
		if idx == 2 {
			dupY = oy + bh/2
		}
		oy += bh + spacing
	}

	// Right column: stack state (after OP_DUP).
	sx := pad + opcodeW + gap
	sy := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.script_stack_label")+" (after OP_DUP)", image.Pt(sx, sy), th.Color.TextMuted)
	sy += gtx.Dp(unit.Dp(24))

	// Group box around stack area.
	stackH := 3*(bh+spacing) + spacing
	groupBox(gtx, th, image.Pt(sx-4, sy-4), stackW+8, stackH+8)

	stackItems := []string{
		i18n.T("diagram.script_pubkey_label"),
		i18n.T("diagram.script_pubkey_label"),
		i18n.T("diagram.script_sig_label"),
	}

	for _, item := range stackItems {
		shadowBox(gtx, th, item, image.Pt(sx, sy), stackW, bh, th.Color.TipBg)
		sy += bh + spacing
	}

	// Badge marking the duplicated entry.
	badge(gtx, th, "2x", image.Pt(sx+stackW+gtx.Dp(unit.Dp(10)), gtx.Dp(unit.Dp(18))+gtx.Dp(unit.Dp(18))+bh), 10, th.Color.Primary)

	// Arrow from OP_DUP to stack.
	dirArrow(gtx, image.Pt(ox+opcodeW, dupY), image.Pt(sx-4, dupY), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
