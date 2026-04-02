# Open Bitcoin Academy

Interactive Bitcoin education for everyone. Offline-first, single-binary, translatable from day one.

## What is this?

An app that teaches Bitcoin from zero. You read short lessons, look at diagrams, answer quizzes, and run real cryptographic operations on your device. No internet required. No account required. Available in multiple languages.

19 chapters covering everything from "What is money?" to Taproot, Miniscript, FROST threshold signatures, Lightning Network, and the broader Bitcoin ecosystem.

Built with [Go](https://go.dev) and [Gio](https://gioui.org). Targets Linux, Windows, and Android.

## Build

```bash
go build -o oba ./cmd/oba && ./oba
```

## Test

```bash
go test ./...
```

## Languages

English, Spanish, Portuguese (BR), French, Swahili.

## License

[Unlicense](UNLICENSE) (public domain).

## Content Attribution

Lesson content is original writing informed by [Mastering Bitcoin: Programming the Open Blockchain (3rd Edition)](https://github.com/bitcoinbook/bitcoinbook) by Andreas M. Antonopoulos and David A. Harding, licensed CC-BY-SA 4.0. Our content is not a copy or derivative of the book. It is independently written educational material covering the same Bitcoin protocol topics, structured for an interactive app format.

## Font Credits

- [NotoSans](https://fonts.google.com/noto/specimen/Noto+Sans) by Google (SIL Open Font License)
- [JetBrains Mono](https://www.jetbrains.com/lp/mono/) by JetBrains (SIL Open Font License)

## Dependencies

| Dependency | License |
|------------|---------|
| [Gio](https://gioui.org) | Unlicense OR MIT |
| [BurntSushi/toml](https://github.com/BurntSushi/toml) | MIT |
| [btcsuite/btcd](https://github.com/btcsuite/btcd) | ISC |
| [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) | BSD-3-Clause |
