// Package oba provides embedded content and locale filesystems.
// This file lives at the module root so go:embed can access
// the content/ and locales/ directories.
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
