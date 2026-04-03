# Open Bitcoin Academy

An offline-first app that teaches the Bitcoin protocol. Lessons, diagrams, quizzes, and cryptographic exercises running on your device. No internet required. No account required. Available in 6 languages.

13 chapters cover the same topics as [Mastering Bitcoin 3rd Edition](https://github.com/bitcoinbook/bitcoinbook). 6 additional chapters go further into Taproot, Miniscript, PSBTs, threshold signatures (FROST/ChillDKG), Lightning, and the broader ecosystem (RGB, Fedimint, Cashu, Ark).

Public domain. Early stage.

## Source

[Codeberg](https://codeberg.org/OpenBitcoinAcademy/oba) (primary) | [GitHub](https://github.com/OpenBitcoinAcademy/oba) (mirror) | [Nostr](https://gitworkshop.dev/npub1qfcefq43er545mwvg69093almdkqs4wmcxaxjg9ad0rzq834qqfq9rhflh/relay.damus.io/oba) (NIP-34)

### Clone via Nostr

With [ngit](https://gitworkshop.dev/ngit) installed:

    git clone nostr://npub1qfcefq43er545mwvg69093almdkqs4wmcxaxjg9ad0rzq834qqfq9rhflh/oba

Patches and issues can be submitted via nostr relays without a Codeberg or GitHub account.

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
