// Command oba is the Open Bitcoin Academy interactive education app.
package main

import (
	"image"
	"io/fs"
	"log"

	gio_app "gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	oba "github.com/openbitcoinacademy/oba"
	"github.com/openbitcoinacademy/oba/internal/app"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/platform"
	"github.com/openbitcoinacademy/oba/internal/ui/screens"
)

func main() {
	// Embedded filesystems. Sub() strips the directory prefix so paths
	// match what the loaders expect (e.g., "en/ui.toml" not "locales/en/ui.toml").
	localeFS, err := fs.Sub(oba.LocaleFS, "locales")
	if err != nil {
		log.Fatalf("locale fs: %v", err)
	}
	contentFS, err := fs.Sub(oba.ContentFS, "content")
	if err != nil {
		log.Fatalf("content fs: %v", err)
	}

	if err := i18n.Init(localeFS); err != nil {
		log.Fatalf("init i18n: %v", err)
	}

	resolver := content.NewResolver(contentFS, localeFS, "en")
	ch, err := content.LoadChapterFromFS(contentFS, "ch01/ch01.toml", resolver)
	if err != nil {
		log.Fatalf("load chapter: %v", err)
	}

	exercises, err := content.LoadChapterExercises(contentFS, "ch01")
	if err != nil {
		log.Printf("load exercises: %v (continuing without)", err)
	}

	progressPath, err := platform.ProgressPath()
	if err != nil {
		log.Printf("progress path: %v (using in-memory)", err)
		progressPath = ""
	}
	progress := content.LoadProgress(progressPath)

	state := app.NewState(ch, progress, progressPath)
	state.ContentFS = contentFS
	state.LocaleFS = localeFS
	state.Exercises = exercises

	go func() {
		w := &gio_app.Window{}
		w.Option(gio_app.Title("Open Bitcoin Academy"))
		w.Option(gio_app.Size(unit.Dp(420), unit.Dp(740)))

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
				return
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
