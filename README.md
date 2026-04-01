# Open Bitcoin Academy

Interactive Bitcoin education for everyone. Offline-first, single-binary, translatable from day one.

## What is this?

An app that teaches Bitcoin from zero. You read short lessons, look at diagrams, answer quizzes, and run real cryptographic operations on your device. No internet required. No account required. Available in multiple languages.

Built with [Go](https://go.dev) and [Gio](https://gioui.org). Targets Linux, Windows, and Android.

## Build

```bash
go build -o oba ./cmd/oba && ./oba
```

## Test

```bash
go test ./...
```

## Requirements

- Go 1.22+
- Gio build dependencies (see ARCHITECTURE.md Section 18.1)

## License

[Unlicense](UNLICENSE) (public domain).
