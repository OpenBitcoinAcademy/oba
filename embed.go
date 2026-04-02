// Package oba provides embedded filesystems for content, locales, and assets.
// This file lives at the module root so go:embed can access top-level directories.
package oba

import "embed"

// ContentFS holds all English content files (chapter TOML and .md).
//
//go:embed all:content
var ContentFS embed.FS

// LocaleFS holds all locale files (ui.toml and translated .md).
//
//go:embed all:locales
var LocaleFS embed.FS

// FontNotoSans holds the NotoSans-Regular.ttf font (Latin coverage).
//
//go:embed assets/fonts/NotoSans-Regular.ttf
var FontNotoSans []byte

// FontJetBrainsMono holds the JetBrainsMono-Regular.ttf font (code/hex).
//
//go:embed assets/fonts/JetBrainsMono-Regular.ttf
var FontJetBrainsMono []byte
