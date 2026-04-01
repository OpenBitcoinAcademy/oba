// Command oba is the Open Bitcoin Academy interactive education app.
package main

import (
	"image"
	"log"
	"os"

	gio_app "gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/app"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/platform"
	"github.com/openbitcoinacademy/oba/internal/ui/screens"
)

func main() {
	localeFS := os.DirFS("locales")
	if err := i18n.Init(localeFS); err != nil {
		log.Fatalf("init i18n: %v", err)
	}

	contentFS := os.DirFS("content")
	resolver := content.NewResolver(contentFS, nil, "en")
	ch, err := content.LoadChapterFromFS(contentFS, "ch01/ch01.toml", resolver)
	if err != nil {
		log.Fatalf("load chapter: %v", err)
	}

	progressPath, err := platform.ProgressPath()
	if err != nil {
		log.Printf("progress path: %v (using in-memory)", err)
		progressPath = ""
	}
	progress := content.LoadProgress(progressPath)

	state := app.NewState(ch, progress, progressPath)

	go func() {
		w := &gio_app.Window{}
		w.Option(gio_app.Title("Open Bitcoin Academy"))
		w.Option(gio_app.Size(unit.Dp(420), unit.Dp(740)))

		// Wire invalidation so navigation triggers redraws.
		state.Invalidate = w.Invalidate

		home := screens.NewHome(state)
		lesson := screens.NewLesson(state)
		settings := screens.NewSettings(state)

		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case gio_app.DestroyEvent:
				if e.Err != nil {
					log.Fatal(e.Err)
				}
				os.Exit(0)
			case gio_app.FrameEvent:
				gtx := gio_app.NewContext(&ops, e)

				paint.FillShape(gtx.Ops, state.Theme.Color.Background,
					clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)}.Op())

				switch state.CurrentScreen {
				case app.ScreenLesson:
					lesson.Layout(gtx)
				case app.ScreenSettings:
					settings.Layout(gtx)
				default:
					home.Layout(gtx)
				}

				e.Frame(gtx.Ops)
			}
		}
	}()
	gio_app.Main()
}
