// Command oba is the Open Bitcoin Academy interactive education app.
package main

import (
	"image"
	"image/color"
	"io/fs"
	"log"
	"os"
	"time"

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
	"github.com/openbitcoinacademy/oba/internal/ui/diagrams"
	"github.com/openbitcoinacademy/oba/internal/ui/screens"
)

func main() {
	// Wire diagram validation so unknown IDs fail at startup.
	content.ValidateDiagramID = diagrams.Validate

	// Embedded filesystems.
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

	locale := detectLocale()
	i18n.SetLocale(locale)

	// Load all chapters from chapters.toml.
	resolver := content.NewResolver(contentFS, localeFS, locale)
	chapters, err := content.LoadAllChapters(contentFS, "chapters.toml", resolver)
	if err != nil {
		log.Fatalf("load chapters: %v", err)
	}

	// Load exercises for all chapters.
	allExercises := make(map[string]*content.ExerciseConfig)
	for _, ch := range chapters {
		exs, err := content.LoadChapterExercises(contentFS, ch.ID)
		if err != nil {
			log.Printf("load exercises for %s: %v (continuing without)", ch.ID, err)
			continue
		}
		for k, v := range exs {
			allExercises[k] = v
		}
	}

	progressPath, err := platform.ProgressPath()
	if err != nil {
		log.Printf("progress path: %v (using in-memory)", err)
		progressPath = ""
	}
	progress := content.LoadProgress(progressPath)

	state := app.NewStateWithFonts(chapters, progress, progressPath, oba.FontNotoSans, oba.FontJetBrainsMono)
	state.ContentFS = contentFS
	state.LocaleFS = localeFS
	state.Exercises = allExercises

	go func() {
		w := &gio_app.Window{}
		w.Option(gio_app.Title("Open Bitcoin Academy"))
		w.Option(gio_app.Size(unit.Dp(420), unit.Dp(740)))
		w.Option(gio_app.StatusColor(state.Theme.Color.Background))
		w.Option(gio_app.NavigationColor(state.Theme.Color.Background))

		state.Invalidate = w.Invalidate

		splash := screens.NewSplash(state, oba.IconPNG)
		home := screens.NewHome(state)
		chapter := screens.NewChapter(state)
		lesson := screens.NewLesson(state)
		settings := screens.NewSettings(state)

		var ops op.Ops
		var lastBarColor color.NRGBA
		var appFadeStart time.Time
		const appFadeDuration = 500 * time.Millisecond
		for {
			switch e := w.Event().(type) {
			case gio_app.DestroyEvent:
				if e.Err != nil {
					log.Fatal(e.Err)
				}
				return
			case gio_app.FrameEvent:
				gtx := gio_app.NewContext(&ops, e)

				// Keep system bars in sync with theme background.
				if bg := state.Theme.Color.Background; bg != lastBarColor {
					w.Option(gio_app.StatusColor(bg))
					w.Option(gio_app.NavigationColor(bg))
					lastBarColor = bg
				}

				paint.FillShape(gtx.Ops, state.Theme.Color.Background,
					clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)}.Op())

				if !splash.Done() {
					splash.Layout(gtx)
					e.Frame(gtx.Ops)
					continue
				}

				// Track when the app content first appears.
				if appFadeStart.IsZero() {
					appFadeStart = time.Now()
				}

				switch state.CurrentScreen {
				case app.ScreenChapter:
					chapter.Layout(gtx)
				case app.ScreenLesson:
					lesson.Layout(gtx)
				case app.ScreenSettings:
					settings.Layout(gtx)
				default:
					home.Layout(gtx)
				}

				// Fade-in overlay: navy fades from opaque to transparent.
				if elapsed := time.Since(appFadeStart); elapsed < appFadeDuration {
					fade := float64(elapsed) / float64(appFadeDuration)
					overlay := state.Theme.Color.Background
					overlay.A = uint8((1 - fade) * 255)
					paint.FillShape(gtx.Ops, overlay,
						clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)}.Op())
					gtx.Execute(op.InvalidateCmd{})
				}

				e.Frame(gtx.Ops)
			}
		}
	}()
	gio_app.Main()
}

// detectLocale checks system environment for a supported locale.
func detectLocale() string {
	available := i18n.Available()
	isAvailable := func(code string) bool {
		for _, a := range available {
			if a == code {
				return true
			}
		}
		return false
	}

	for _, env := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		val := os.Getenv(env)
		if val == "" || val == "C" || val == "POSIX" {
			continue
		}
		code := val
		for i, c := range code {
			if c == '_' || c == '.' || c == '@' {
				code = code[:i]
				break
			}
		}
		if isAvailable(code) {
			return code
		}
	}

	return "en"
}
