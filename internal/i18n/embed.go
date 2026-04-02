// Embed strategy: the real go:embed directives live in embed.go at
// the module root (package oba). The embedded fs.FS is passed to
// Init() from cmd/oba/main.go via fs.Sub() to strip the "locales/"
// prefix. This file documents the pattern for future reference.
package i18n
