// Command oba is the Open Bitcoin Academy interactive education app.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/openbitcoinacademy/oba/internal/i18n"
)

//go:generate echo "Placeholder: go:embed will be added when content exists"

func main() {
	localeFS := os.DirFS("locales")
	if err := i18n.Init(localeFS); err != nil {
		log.Fatalf("init i18n: %v", err)
	}

	fmt.Println(i18n.T("chapters.ch01_title"))

	locales := i18n.Available()
	fmt.Printf("Available locales: %v\n", locales)

	_ = fs.FS(localeFS) // Silence unused import in future.
}
