# Open Bitcoin Academy

Bitcoin education matters more than anything else in Bitcoin. This app teaches the protocol from zero, with no shortcuts.

## What

An offline-first app covering Bitcoin from "What is money?" through Taproot, Miniscript, FROST threshold signatures, Lightning, and the ecosystem. 19 chapters aligned 1:1 with [Mastering Bitcoin 3rd Edition](https://github.com/bitcoinbook/bitcoinbook). Short lessons, diagrams, quizzes, and real cryptographic operations running on your device. No internet required. No account required.

Built with [Go](https://go.dev) and [Gio](https://gioui.org). Available in English, Spanish, Portuguese, French, and Swahili.

This is the initial release. We build from here.

## Source

[Codeberg](https://codeberg.org/OpenBitcoinAcademy/oba) (primary) | [GitHub](https://github.com/OpenBitcoinAcademy/oba) (mirror)

## Build

```bash
go build -o oba ./cmd/oba && ./oba
```

## Test

```bash
go test ./...
```

## Roadmap

- Better UX with game-like interactivity
- Audio lessons
- Lexicon mode (look up what you need) alongside the structured course
- More languages
- Run on any device you own (Android, iOS, Windows, macOS, Linux)
- Accessibility

## Contributing

Translations and code contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

**Translate**: copy `locales/en/`, translate the files, run `go run ./tools/translate-check`. You never touch Go code or TOML structure.

**Code**: branch from `main`, write tests, use conventional commits. Open a PR on [Codeberg](https://codeberg.org/OpenBitcoinAcademy/oba).

## License

[Unlicense](UNLICENSE) (public domain).

## Content Attribution

Lesson content is original writing informed by [Mastering Bitcoin: Programming the Open Blockchain (3rd Edition)](https://github.com/bitcoinbook/bitcoinbook) by Andreas M. Antonopoulos and David A. Harding, licensed CC-BY-SA 4.0. Our content is not a copy or derivative of the book. It is independently written educational material covering the same Bitcoin protocol topics, structured for an interactive app format.

## Font Credits

- [NotoSans](https://fonts.google.com/noto/specimen/Noto+Sans) by Google (SIL Open Font License)
- [JetBrains Mono](https://www.jetbrains.com/lp/mono/) by JetBrains (SIL Open Font License)

## Contact

[npub1qfcefq43er545mwvg69093almdkqs4wmcxaxjg9ad0rzq834qqfq9rhflh](https://njump.me/npub1qfcefq43er545mwvg69093almdkqs4wmcxaxjg9ad0rzq834qqfq9rhflh) on Nostr

## Dependencies

| Dependency | License |
|------------|---------|
| [Gio](https://gioui.org) | Unlicense OR MIT |
| [BurntSushi/toml](https://github.com/BurntSushi/toml) | MIT |
| [btcsuite/btcd](https://github.com/btcsuite/btcd) | ISC |
| [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) | BSD-3-Clause |
