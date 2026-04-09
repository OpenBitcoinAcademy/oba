// Package i18n provides internationalization for UI chrome strings.
// The real go:embed directives live in embed.go at the module root
// (package oba). The embedded fs.FS is passed to Init() from
// cmd/oba/main.go via fs.Sub() to strip the "locales/" prefix.
package i18n
