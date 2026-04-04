package screens

import (
	"bytes"
	"image"
	"image/png"
	"time"

	"gioui.org/f32"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/app"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Splash shows the logo and tagline with a fade-in animation.
type Splash struct {
	state   *app.State
	start   time.Time
	iconOp  paint.ImageOp
	iconSet bool
	done    bool
}

// NewSplash creates a splash screen.
func NewSplash(state *app.State, iconData []byte) *Splash {
	s := &Splash{state: state, start: time.Now()}
	if img, err := png.Decode(bytes.NewReader(iconData)); err == nil {
		s.iconOp = paint.NewImageOp(img)
		s.iconSet = true
	}
	return s
}

// Done reports whether the splash animation is complete.
func (s *Splash) Done() bool {
	return s.done
}

// Layout renders the splash screen.
func (s *Splash) Layout(gtx layout.Context) layout.Dimensions {
	th := s.state.Theme
	elapsed := time.Since(s.start)

	// Timeline: fade in 0-600ms, hold 600-2500ms, fade out 2500-3000ms.
	const fadeIn = 600 * time.Millisecond
	const holdEnd = 2500 * time.Millisecond
	const total = 3 * time.Second

	// Always request next frame until done.
	gtx.Execute(op.InvalidateCmd{})

	var alpha uint8
	if elapsed < fadeIn {
		alpha = uint8(float64(elapsed) / float64(fadeIn) * 255)
	} else if elapsed < holdEnd {
		alpha = 255
	} else if elapsed < total {
		fade := float64(elapsed-holdEnd) / float64(total-holdEnd)
		alpha = uint8((1 - fade) * 255)
	} else {
		alpha = 0
		s.done = true
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(gtx,
			// Logo.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !s.iconSet {
					return layout.Dimensions{}
				}
				return layout.Inset{Bottom: unit.Dp(32)}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return s.drawIcon(gtx, alpha)
					},
				)
			}),
			// "Learn Bitcoin"
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return s.drawLabel(gtx, th, "Learn Bitcoin.", th.Text.H1, alpha)
			}),
			// "Build Understanding."
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: unit.Dp(4)}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return s.drawLabel(gtx, th, "Build Understanding.", th.Text.H2, alpha)
					},
				)
			}),
		)
	})
}

func (s *Splash) drawIcon(gtx layout.Context, alpha uint8) layout.Dimensions {
	iconSize := gtx.Dp(unit.Dp(200))
	sz := s.iconOp.Size()
	if sz.X == 0 || sz.Y == 0 {
		return layout.Dimensions{}
	}

	scale := float32(iconSize) / float32(sz.X)
	defer op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(scale, scale))).Push(gtx.Ops).Pop()

	s.iconOp.Filter = paint.FilterLinear
	s.iconOp.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	outSize := image.Pt(iconSize, iconSize)
	return layout.Dimensions{Size: outSize}
}

func (s *Splash) drawLabel(gtx layout.Context, th *theme.Theme, label string, size unit.Sp, alpha uint8) layout.Dimensions {
	lbl := material.Label(th.Material, size, label)
	lbl.Color = th.Color.Text
	lbl.Color.A = alpha
	lbl.Alignment = text.Middle
	lbl.Font.Weight = font.Bold
	return lbl.Layout(gtx)
}
