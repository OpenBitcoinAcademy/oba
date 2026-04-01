# Contributing to Open Bitcoin Academy

## Development Environment

1. Install Go 1.22+ from [go.dev](https://go.dev/dl/)
2. Install Gio build dependencies (see ARCHITECTURE.md Section 18.1)
3. Install linters: `staticcheck`, `golangci-lint`
4. Clone the repo and verify: `make test && make lint`

## Adding a Translation

1. Create a locale directory: `locales/{lang}/`
2. Copy `locales/en/ui.toml` and translate all values
3. Mirror all `.md` files from `content/` into `locales/{lang}/`
4. Run `make translate-check` to verify completeness
5. You never need to touch TOML structure files or Go code

## Writing Lesson Content

- Use the supported Markdown subset (see ARCHITECTURE.md Section 6)
- Follow the writing style guide in CLAUDE.md
- Quiz format: first line is the question, `- ` lines are options
- Correct answer indices live in TOML, not in the `.md` file

## Submitting Changes

1. Branch from `main`: `git checkout -b feature/your-feature`
2. Write code and tests
3. Run `make test && make lint`
4. Commit with conventional commit messages: `feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `ci:`
5. Open a PR, squash-merge to main

## Reporting Issues

Use GitHub Issues. Include: steps to reproduce, expected behavior, actual behavior, platform/OS.
