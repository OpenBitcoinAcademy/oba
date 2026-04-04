# Contributing to Open Bitcoin Academy

## Adding a Translation

1. Create a locale directory: `locales/{lang}/`
2. Copy `locales/en/ui.toml` and translate all values (button labels, chapter titles, UI text)
3. Mirror all `.md` files from `content/` into `locales/{lang}/` and translate them
4. Run `go run ./tools/translate-check` to verify completeness
5. You translate one `.toml` file and all `.md` files. No Go code, no structure changes

## Writing Lesson Content

Supported Markdown: paragraphs, `**bold**`, `*italic*`, `` `inline code` ``, fenced code blocks, `##`/`###` headings, `$...$` inline math. No images, tables, HTML, or links.

Quiz format: first line is the question, lines starting with `- ` are options. Correct answer index lives in TOML, not in the `.md` file.

## Submitting Changes

1. Branch from `main`: `git checkout -b feature/your-feature`
2. Write code and tests
3. Run `go test ./...` and `go vet ./...`
4. Commit with conventional prefixes: `feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `ci:`
5. Open a PR on [Codeberg](https://codeberg.org/OpenBitcoinAcademy/oba), squash-merge to main

## Reporting Issues

Use [Codeberg Issues](https://codeberg.org/OpenBitcoinAcademy/oba/issues). Include: steps to reproduce, expected behavior, actual behavior, platform/OS.

## Contributing via Nostr

With [ngit](https://gitworkshop.dev/ngit) installed, you can submit patches and file issues via nostr relays. No account on any platform required. See the repo on [gitworkshop.dev](https://gitworkshop.dev).

