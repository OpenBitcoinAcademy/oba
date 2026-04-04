package diagrams

import (
	"encoding/hex"
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/bitcoin"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// HashFlow draws Input -> [SHA-256] -> Hash with a real computed hash.
type HashFlow struct{}

func (d *HashFlow) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(90))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	// No background fill: inherits page background.

	// Responsive: input 20%, process 22%, output 36% of width, rest is gaps.
	inputW := pct(w, 20)
	procW := pct(w, 22)
	outputW := pct(w, 36)
	bh := gtx.Dp(unit.Dp(36))

	usable := w - 2*pad
	gapTotal := usable - inputW - procW - outputW
	gap := gapTotal / 2
	if gap < 8 {
		gap = 8
	}

	y := (h - bh) / 2

	x1 := pad
	x2 := x1 + inputW + gap
	x3 := x2 + procW + gap

	// Compute real hash for the example input.
	input := "Hello"
	hashBytes := bitcoin.SHA256([]byte(input))
	hashHex := hex.EncodeToString(hashBytes[:8]) + "..."

	shadowBox(gtx, th, "\""+input+"\"", image.Pt(x1, y), inputW, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x1+inputW, y+bh/2), image.Pt(x2, y+bh/2), 1.8, withAlpha(th.Color.TextMuted, 160))
	shadowBox(gtx, th, i18n.T("diagram.sha256"), image.Pt(x2, y), procW, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x2+procW, y+bh/2), image.Pt(x3, y+bh/2), 1.8, withAlpha(th.Color.TextMuted, 160))
	shadowBox(gtx, th, hashHex, image.Pt(x3, y), outputW, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
