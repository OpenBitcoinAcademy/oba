// Package theme defines the OBA visual design: colors, spacing, and
// typography. Optimized for readability on cheap phones with small
// screens in direct sunlight (high contrast, large touch targets).
package theme

import (
	"image/color"
	"log"
	"os"
	"strings"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
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
// Clean white base, navy text (matching logo), muted academic accents.
func LightPalette() Colors {
	return Colors{
		Background: rgb(0xFFFFFF),
		Surface:    rgb(0xF5F5F8),
		Primary:    rgb(0xF7931A), // Bitcoin orange
		OnPrimary:  rgb(0xFFFFFF),
		Text:       rgb(0x181C46), // Navy (logo)
		TextMuted:  rgb(0x6B6E82),
		Accent:     rgb(0x4A5BA0),
		Error:      rgb(0xC83232),
		Success:    rgb(0x3A7D42),
		Warning:    rgb(0xD97515),
		InfoBg:     rgba(0x181C46, 0x0F), // Subtle navy tint
		WarningBg:  rgba(0x181C46, 0x0F),
		TipBg:      rgba(0x181C46, 0x0F),
		Divider:    rgb(0xD8D9E0),
	}
}

// DarkPalette returns a high-contrast dark theme.
// Navy base (logo background), cool grays, Bitcoin orange accents.
func DarkPalette() Colors {
	return Colors{
		Background: rgb(0x181C46), // Logo navy
		Surface:    rgb(0x1F2350),
		Primary:    rgb(0xF7931A), // Bitcoin orange stays
		OnPrimary:  rgb(0x181C46), // Navy on orange
		Text:       rgb(0xE8E8EC),
		TextMuted:  rgb(0x8A8DA0),
		Accent:     rgb(0x7B8CDE),
		Error:      rgb(0xE05555),
		Success:    rgb(0x5CA065),
		Warning:    rgb(0xE09040),
		InfoBg:     rgba(0xFFFFFF, 0x14), // Transparent overlay
		WarningBg:  rgba(0xFFFFFF, 0x14),
		TipBg:      rgba(0xFFFFFF, 0x14),
		Divider:    rgb(0x2A2E58),
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

// New creates a Theme using the detected system preference and default fonts.
func New() *Theme {
	mode := DetectSystemMode()
	return NewWithFonts(mode, nil, nil)
}

// NewWithFonts creates a Theme with custom font data.
// If notoSans or jetBrainsMono are nil, Gio's built-in fonts are used.
func NewWithFonts(mode Mode, notoSans, jetBrainsMono []byte) *Theme {
	var faces []font.FontFace

	if len(notoSans) > 0 {
		face, err := opentype.Parse(notoSans)
		if err != nil {
			log.Printf("parse NotoSans: %v (using default)", err)
		} else {
			faces = append(faces, font.FontFace{
				Font: font.Font{Typeface: "NotoSans"},
				Face: face,
			})
		}
	}

	if len(jetBrainsMono) > 0 {
		face, err := opentype.Parse(jetBrainsMono)
		if err != nil {
			log.Printf("parse JetBrainsMono: %v (using default)", err)
		} else {
			faces = append(faces, font.FontFace{
				Font: font.Font{Typeface: "monospace"},
				Face: face,
			})
		}
	}

	// Start with Gio's built-in Go fonts (includes bold, italic variants),
	// then prepend our custom fonts so they take priority.
	collection := gofont.Collection()
	if len(faces) > 0 {
		collection = append(faces, collection...)
	}

	mat := material.NewTheme()
	mat.Shaper = text.NewShaper(text.WithCollection(collection))

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

func rgba(c uint32, a uint8) color.NRGBA {
	return color.NRGBA{
		R: uint8(c >> 16),
		G: uint8(c >> 8),
		B: uint8(c),
		A: a,
	}
}
