// Command oba is the Open Bitcoin Academy interactive education app.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
)

func main() {
	localeFS := os.DirFS("locales")
	if err := i18n.Init(localeFS); err != nil {
		log.Fatalf("init i18n: %v", err)
	}

	// Load Chapter 1.
	contentFS := os.DirFS("content")
	resolver := content.NewResolver(contentFS, nil, "en")

	ch, err := content.LoadChapterFromFS(contentFS, "ch01/ch01.toml", resolver)
	if err != nil {
		log.Fatalf("load chapter: %v", err)
	}

	fmt.Printf("Chapter: %s\n", i18n.T(ch.TitleKey))
	fmt.Printf("  %s\n\n", i18n.T(ch.DescKey))

	for _, lesson := range ch.Lessons {
		fmt.Printf("  Lesson: %s\n", i18n.T(lesson.TitleKey))
		fmt.Printf("    Sections: %d\n", len(lesson.Sections))
		for j, sec := range lesson.Sections {
			switch s := sec.(type) {
			case *content.TextSection:
				// Show first 60 chars of content.
				preview := s.Content
				if len(preview) > 60 {
					preview = preview[:60] + "..."
				}
				fmt.Printf("      [%d] text: %s\n", j+1, preview)
			case *content.CalloutSection:
				fmt.Printf("      [%d] callout (%s)\n", j+1, s.Style)
			case *content.DiagramSection:
				fmt.Printf("      [%d] diagram: %s\n", j+1, s.DiagramID)
			case *content.QuizSection:
				fmt.Printf("      [%d] quiz: %s (%d options)\n", j+1, s.Question, len(s.Options))
			case *content.InteractiveSection:
				fmt.Printf("      [%d] exercise: %s\n", j+1, s.ExerciseID)
			}
		}
		fmt.Println()
	}

	fmt.Printf("Locales: %v\n", i18n.Available())
}
