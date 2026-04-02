package screens

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/app"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	obamath "github.com/openbitcoinacademy/oba/internal/math"
	"github.com/openbitcoinacademy/oba/internal/ui/components"
	"github.com/openbitcoinacademy/oba/internal/ui/diagrams"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ExerciseWidget is any exercise that can lay itself out.
type ExerciseWidget interface {
	Layout(gtx layout.Context) layout.Dimensions
}

// Lesson displays a scrollable lesson with all its sections.
type Lesson struct {
	state      *app.State
	list       widget.List
	backBtn    widget.Clickable
	quizzes    map[int]*components.Quiz
	exercises  map[int]ExerciseWidget
	lastLesson int // tracks which lesson we last rendered
}

// NewLesson creates the lesson screen.
func NewLesson(state *app.State) *Lesson {
	l := &Lesson{
		state:      state,
		quizzes:    make(map[int]*components.Quiz),
		exercises:  make(map[int]ExerciseWidget),
		lastLesson: -1,
	}
	l.list.Axis = layout.Vertical
	return l
}

// reset clears all per-lesson widget state.
func (l *Lesson) reset() {
	l.quizzes = make(map[int]*components.Quiz)
	l.exercises = make(map[int]ExerciseWidget)
	l.list.Position = layout.Position{}
	l.lastLesson = l.state.CurrentLesson
}

// Layout renders the lesson screen.
func (l *Lesson) Layout(gtx layout.Context) layout.Dimensions {
	// Reset widget state when switching lessons.
	if l.state.LessonDirty || l.lastLesson != l.state.CurrentLesson {
		l.reset()
		l.state.LessonDirty = false
	}

	th := l.state.Theme
	lesson := l.state.Chapter.Lessons[l.state.CurrentLesson]

	if l.backBtn.Clicked(gtx) {
		l.state.NavigateChapterList()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Top bar.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top: th.Space.Small, Bottom: th.Space.Small,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &l.backBtn, i18n.T("nav.previous"))
						btn.Background = th.Color.Surface
						btn.Color = th.Color.Primary
						return btn.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, i18n.T(lesson.TitleKey))
						lbl.Color = th.Color.Text
						return lbl.Layout(gtx)
					}),
				)
			})
		}),
		// Divider.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, th.Color.Divider,
				clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, 1)}.Op())
			return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, 1)}
		}),
		// Scrollable sections.
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.List(th.Material, &l.list).Layout(gtx, len(lesson.Sections),
				func(gtx layout.Context, i int) layout.Dimensions {
					return layout.Inset{
						Left: th.Space.Medium, Right: th.Space.Medium,
						Top: th.Space.Medium,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return l.renderSection(gtx, lesson.Sections[i], i)
					})
				},
			)
		}),
	)
}

func (l *Lesson) renderSection(gtx layout.Context, sec content.Section, idx int) layout.Dimensions {
	th := l.state.Theme

	switch s := sec.(type) {
	case *content.TextSection:
		md := &components.Markdown{Content: s.Content, Theme: th}
		return md.Layout(gtx)

	case *content.CalloutSection:
		co := &components.Callout{Content: s.Content, Style: s.Style, Theme: th}
		return co.Layout(gtx)

	case *content.DiagramSection:
		if dw, ok := diagrams.Registry[s.DiagramID]; ok {
			return dw.Layout(gtx, th)
		}
		return placeholderSection(gtx, th, "Diagram: "+s.DiagramID)

	case *content.QuizSection:
		q, ok := l.quizzes[idx]
		if !ok {
			q = components.NewQuiz(s, th)
			l.quizzes[idx] = q
		}
		return q.Layout(gtx)

	case *content.InteractiveSection:
		ex, ok := l.exercises[idx]
		if !ok {
			ex = l.createExercise(s.ExerciseID)
			l.exercises[idx] = ex
		}
		if ex != nil {
			return ex.Layout(gtx)
		}
		return placeholderSection(gtx, th, "Exercise: "+s.ExerciseID)

	case *content.FormulaSection:
		renderer := obamath.NewRenderer(th.Material, th.Text.Body)
		renderer.Color = th.Color.Text
		return layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(8)}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min = image.Point{}
				return renderer.Layout(gtx, s.Formula)
			})

	default:
		return placeholderSection(gtx, th, "Unknown section")
	}
}

func (l *Lesson) createExercise(id string) ExerciseWidget {
	th := l.state.Theme
	cfg := l.state.Exercises[id] // nil if not loaded, widgets handle nil

	// Dispatch by exercise type from config, fall back to ID matching.
	exType := id
	if cfg != nil {
		exType = cfg.Type
	}

	switch exType {
	case "hash_explorer", "ex01_hash":
		return components.NewHashExplorer(th, cfg)
	case "key_generator", "ex02_keys":
		return components.NewKeyGenerator(th, cfg)
	case "address_builder", "ex03_address":
		return components.NewAddressBuilder(th, cfg)
	case "script_debugger", "ex04_script_debugger":
		return components.NewScriptDebugger(th, cfg)
	case "tx_inspector", "ex05_tx_inspector":
		return components.NewTxInspector(th, cfg)
	case "tx_builder", "ex06_tx_builder":
		return components.NewTxBuilder(th, cfg)
	case "block_inspector", "ex07_block_inspector":
		return components.NewBlockInspector(th, cfg)
	case "mining_simulator", "ex08_mining_simulator":
		return components.NewMiningSimulator(th, cfg)
	default:
		return nil
	}
}

func placeholderSection(gtx layout.Context, th *theme.Theme, text string) layout.Dimensions {
	return layout.Inset{
		Top: unit.Dp(8), Bottom: unit.Dp(8),
		Left: unit.Dp(12), Right: unit.Dp(12),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		paint.FillShape(gtx.Ops, th.Color.Divider,
			clip.Rect{Max: gtx.Constraints.Max}.Op())
		return layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(8),
			Left: unit.Dp(8)}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, text)
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			})
	})
}
