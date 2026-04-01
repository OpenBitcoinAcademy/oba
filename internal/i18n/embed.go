package i18n

import "embed"

// LocaleFS holds all locale files. The embed path is relative to the
// module root, so this file must be generated or the embed must be
// injected from cmd/oba/main.go. For now, the Init function accepts
// any fs.FS, and the actual embedding happens in cmd/oba.
//
// This placeholder exists to document the intended embed pattern.
// The real go:embed directive lives in cmd/oba/main.go because
// go:embed can only access files in the same package directory or below.

// PlaceholderFS is unused. The real FS is passed via Init().
var PlaceholderFS embed.FS
