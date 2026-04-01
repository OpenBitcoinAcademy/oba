// Package theme defines the OBA visual design: colors, spacing, and
// typography. Optimized for readability on cheap phones with small
// screens in direct sunlight (high contrast, large touch targets).
package theme

import (
	"image/color"
	"os"
	"strings"

	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Mode identifies the color scheme.
type Mode int

const (
	ModeLight Mode = iota
	ModeDark
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

// DarkPalette returns a high-contrast dark theme.
func DarkPalette() Colors {
	return Colors{
		Background: rgb(0x121212),
		Surface:    rgb(0x1E1E1E),
		Primary:    rgb(0xF7931A), // Bitcoin orange stays
		OnPrimary:  rgb(0x000000),
		Text:       rgb(0xE0E0E0),
		TextMuted:  rgb(0x9E9E9E),
		Accent:     rgb(0x64B5F6),
		Error:      rgb(0xEF5350),
		Success:    rgb(0x66BB6A),
		Warning:    rgb(0xFFA726),
		InfoBg:     rgb(0x1A237E),
		WarningBg:  rgb(0x3E2723),
		TipBg:      rgb(0x1B5E20),
		Divider:    rgb(0x333333),
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
	Mode     Mode
}

// New creates a Theme using the detected system preference.
func New() *Theme {
	mode := DetectSystemMode()
	return NewWithMode(mode)
}

// NewWithMode creates a Theme with the specified color mode.
func NewWithMode(mode Mode) *Theme {
	mat := material.NewTheme()
	var colors Colors
	if mode == ModeDark {
		colors = DarkPalette()
	} else {
		colors = LightPalette()
	}

	mat.Palette.Bg = colors.Background
	mat.Palette.Fg = colors.Text
	mat.Palette.ContrastBg = colors.Primary
	mat.Palette.ContrastFg = colors.OnPrimary

	return &Theme{
		Material: mat,
		Color:    colors,
		Space:    DefaultSpacing(),
		Text:     DefaultTextSizes(),
		Mode:     mode,
	}
}

// SetMode switches the theme between light and dark.
func (t *Theme) SetMode(mode Mode) {
	t.Mode = mode
	if mode == ModeDark {
		t.Color = DarkPalette()
	} else {
		t.Color = LightPalette()
	}
	t.Material.Palette.Bg = t.Color.Background
	t.Material.Palette.Fg = t.Color.Text
	t.Material.Palette.ContrastBg = t.Color.Primary
	t.Material.Palette.ContrastFg = t.Color.OnPrimary
}

// DetectSystemMode checks environment for dark mode preference.
// Checks OBA_THEME env var first, then common desktop indicators.
func DetectSystemMode() Mode {
	// Explicit override.
	if v := os.Getenv("OBA_THEME"); v != "" {
		if strings.EqualFold(v, "dark") {
			return ModeDark
		}
		return ModeLight
	}

	// GTK theme hint (GNOME, XFCE, etc).
	if v := os.Getenv("GTK_THEME"); strings.Contains(strings.ToLower(v), "dark") {
		return ModeDark
	}

	// Qt/KDE hint.
	if v := os.Getenv("QT_STYLE_OVERRIDE"); strings.Contains(strings.ToLower(v), "dark") {
		return ModeDark
	}

	// Freedesktop color-scheme preference.
	if v := os.Getenv("XDG_CURRENT_DESKTOP"); v != "" {
		// Could query dbus org.freedesktop.appearance color-scheme,
		// but that requires cgo or a dbus library. For now, env vars suffice.
	}

	return ModeLight
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
