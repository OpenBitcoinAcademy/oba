package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// HDWalletTree draws an HD wallet key derivation tree for Ch05.
// Seed -> Master Key -> Account 0, Account 1.
// Under Account 0: Address 0, Address 1.
type HDWalletTree struct{}

func (d *HDWalletTree) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bw := pct(w, 26)
	smallW := pct(w, 22)
	bh := gtx.Dp(unit.Dp(30))
	rowGap := gtx.Dp(unit.Dp(14))
	lc := th.Color.TextMuted
	lw := float32(1.5)

	// Row 0: Seed (centered).
	seedW := pct(w, 22)
	seedX := (w - seedW) / 2
	seedY := gtx.Dp(unit.Dp(4))
	box(gtx, th, i18n.T("diagram.seed"), image.Pt(seedX, seedY), seedW, bh, th.Color.WarningBg)

	// Row 1: Master Key (centered).
	masterY := seedY + bh + rowGap
	masterX := (w - bw) / 2
	box(gtx, th, i18n.T("diagram.master_key"), image.Pt(masterX, masterY), bw, bh, th.Color.Primary)

	// Line: Seed -> Master Key.
	line(gtx, image.Pt(seedX+seedW/2, seedY+bh), image.Pt(masterX+bw/2, masterY), lw, lc)

	// Row 2: Account 0 (left), Account 1 (right).
	acctY := masterY + bh + rowGap
	usable := w - 2*pad
	acctGap := usable / 5
	acct0X := pad + acctGap/2
	acct1X := w - pad - acctGap/2 - bw

	box(gtx, th, i18n.T("diagram.account")+" 0", image.Pt(acct0X, acctY), bw, bh, th.Color.InfoBg)
	box(gtx, th, i18n.T("diagram.account")+" 1", image.Pt(acct1X, acctY), bw, bh, th.Color.InfoBg)

	// Lines: Master Key -> Account 0, Account 1.
	masterMidX := masterX + bw/2
	line(gtx, image.Pt(masterMidX, masterY+bh), image.Pt(acct0X+bw/2, acctY), lw, lc)
	line(gtx, image.Pt(masterMidX, masterY+bh), image.Pt(acct1X+bw/2, acctY), lw, lc)

	// Row 3: Address 0 and Address 1 under Account 0.
	addrY := acctY + bh + rowGap
	addrGap := pct(w, 3)
	addr0X := acct0X - smallW/4
	addr1X := addr0X + smallW + addrGap

	box(gtx, th, i18n.T("diagram.hd_address")+" 0", image.Pt(addr0X, addrY), smallW, bh, th.Color.TipBg)
	box(gtx, th, i18n.T("diagram.hd_address")+" 1", image.Pt(addr1X, addrY), smallW, bh, th.Color.TipBg)

	// Lines: Account 0 -> Address 0, Address 1.
	acct0MidX := acct0X + bw/2
	line(gtx, image.Pt(acct0MidX, acctY+bh), image.Pt(addr0X+smallW/2, addrY), lw, lc)
	line(gtx, image.Pt(acct0MidX, acctY+bh), image.Pt(addr1X+smallW/2, addrY), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
