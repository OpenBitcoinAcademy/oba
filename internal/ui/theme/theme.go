// Package theme defines the OBA visual design: colors, spacing, and
// typography. Optimized for readability on cheap phones with small
// screens in direct sunlight (high contrast, large touch targets).
package theme

import (
	"image/color"

	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Colors defines the OBA color palette. High contrast for readability.
type Colors struct {
	Background color.NRGBA
	Surface    color.NRGBA
	Primary    color.NRGBA
	OnPrimary  color.NRGBA
	Text       color.NRGBA
	TextMuted  color.NRGBA
	Accent     color.NRGBA
	Error      color.NRGBA
	Success    color.NRGBA
	Warning    color.NRGBA
	InfoBg     color.NRGBA
	WarningBg  color.NRGBA
	TipBg      color.NRGBA
	Divider    color.NRGBA
}

// LightPalette returns the default high-contrast light theme.
func LightPalette() Colors {
	return Colors{
		Background: rgb(0xFAFAFA),
		Surface:    rgb(0xFFFFFF),
		Primary:    rgb(0xF7931A), // Bitcoin orange
		OnPrimary:  rgb(0xFFFFFF),
		Text:       rgb(0x1A1A1A),
		TextMuted:  rgb(0x666666),
		Accent:     rgb(0x4A90D9),
		Error:      rgb(0xD32F2F),
		Success:    rgb(0x388E3C),
		Warning:    rgb(0xF57C00),
		InfoBg:     rgb(0xE3F2FD),
		WarningBg:  rgb(0xFFF3E0),
		TipBg:      rgb(0xE8F5E9),
		Divider:    rgb(0xE0E0E0),
	}
}

// Spacing holds standard spacing values in Dp.
type Spacing struct {
	XSmall unit.Dp // 4
	Small  unit.Dp // 8
	Medium unit.Dp // 16
	Large  unit.Dp // 24
	XLarge unit.Dp // 32
}

// DefaultSpacing returns the standard spacing scale.
func DefaultSpacing() Spacing {
	return Spacing{
		XSmall: 4,
		Small:  8,
		Medium: 16,
		Large:  24,
		XLarge: 32,
	}
}

// TextSizes holds font sizes in Sp. Sized for 5-inch 720p screens.
type TextSizes struct {
	Body      unit.Sp // 16
	BodySmall unit.Sp // 14
	H1        unit.Sp // 28
	H2        unit.Sp // 22
	H3        unit.Sp // 18
	Code      unit.Sp // 14
	Caption   unit.Sp // 12
}

// DefaultTextSizes returns standard font sizes.
func DefaultTextSizes() TextSizes {
	return TextSizes{
		Body:      16,
		BodySmall: 14,
		H1:        28,
		H2:        22,
		H3:        18,
		Code:      14,
		Caption:   12,
	}
}

// Theme wraps a Gio material.Theme with OBA-specific design tokens.
type Theme struct {
	Material *material.Theme
	Color    Colors
	Space    Spacing
	Text     TextSizes
}

// New creates a Theme with the light palette and default settings.
func New() *Theme {
	mat := material.NewTheme()
	colors := LightPalette()

	// Apply OBA colors to the material theme.
	mat.Palette.Bg = colors.Background
	mat.Palette.Fg = colors.Text
	mat.Palette.ContrastBg = colors.Primary
	mat.Palette.ContrastFg = colors.OnPrimary

	return &Theme{
		Material: mat,
		Color:    colors,
		Space:    DefaultSpacing(),
		Text:     DefaultTextSizes(),
	}
}

// MinTouchTarget is the minimum touch target size (48dp per Material guidelines).
const MinTouchTarget = unit.Dp(48)

func rgb(c uint32) color.NRGBA {
	return color.NRGBA{
		R: uint8(c >> 16),
		G: uint8(c >> 8),
		B: uint8(c),
		A: 0xFF,
	}
}
