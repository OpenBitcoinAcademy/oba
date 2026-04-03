package components

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Quiz renders a multiple-choice question with radio buttons or checkboxes.
type Quiz struct {
	Section    *content.QuizSection
	Theme      *theme.Theme
	selected   []bool
	checked    bool
	correct    bool
	attempts   int
	revealed   bool // correct answers shown after 3 wrong attempts
	checkBtn   widget.Clickable
	retryBtn   widget.Clickable
	optionBtns []widget.Clickable
}

// NewQuiz creates a quiz widget for the given section.
func NewQuiz(section *content.QuizSection, th *theme.Theme) *Quiz {
	return &Quiz{
		Section:    section,
		Theme:      th,
		selected:   make([]bool, len(section.Options)),
		optionBtns: make([]widget.Clickable, len(section.Options)),
	}
}

// Layout renders the quiz.
func (q *Quiz) Layout(gtx layout.Context) layout.Dimensions {
	// Handle check button click.
	if q.checkBtn.Clicked(gtx) && !q.checked {
		q.checked = true
		q.correct = q.isCorrect()
		if !q.correct {
			q.attempts++
			if q.attempts >= 3 {
				q.revealed = true
			}
		}
	}

	// Handle retry button click.
	if q.retryBtn.Clicked(gtx) {
		q.checked = false
		q.correct = false
		for i := range q.selected {
			q.selected[i] = false
		}
	}

	// Handle option clicks.
	for i := range q.optionBtns {
		if q.optionBtns[i].Clicked(gtx) {
			q.SelectOption(i)
		}
	}

	var children []layout.FlexChild

	// Question.
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(q.Theme.Material, q.Theme.Text.H3, q.Section.Question)
		lbl.Color = q.Theme.Color.Text
		return lbl.Layout(gtx)
	}))
	children = append(children, layout.Rigid(layout.Spacer{Height: q.Theme.Space.Medium}.Layout))

	// Options.
	multiAnswer := len(q.Section.Correct) > 1
	for i := range q.Section.Options {
		idx := i
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return q.renderOption(gtx, idx, multiAnswer)
		}))
	}

	children = append(children, layout.Rigid(layout.Spacer{Height: q.Theme.Space.Medium}.Layout))

	// Check button or result.
	children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		if q.checked {
			return q.renderResult(gtx)
		}
		btn := material.Button(q.Theme.Material, &q.checkBtn, i18n.T("quiz.check"))
		btn.Background = q.Theme.Color.Primary
		btn.Color = q.Theme.Color.OnPrimary
		return btn.Layout(gtx)
	}))

	return layout.Inset{
		Top: q.Theme.Space.Medium, Bottom: q.Theme.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	})
}

func (q *Quiz) renderOption(gtx layout.Context, idx int, multi bool) layout.Dimensions {
	optColor := q.Theme.Color.Text

	if q.checked && q.correct {
		// Answered correctly: highlight correct selections green.
		if q.selected[idx] && q.inCorrect(idx) {
			optColor = q.Theme.Color.Success
		}
	} else if q.checked && !q.correct {
		// Wrong answer.
		if q.selected[idx] {
			optColor = q.Theme.Color.Error
		}
		// Only reveal correct answers after 3 failed attempts.
		if q.revealed && q.inCorrect(idx) {
			optColor = q.Theme.Color.Success
		}
	}

	return layout.Inset{Bottom: unit.Dp(6)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Clickable(gtx, &q.optionBtns[idx], func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				// Indicator (circle or square).
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					size := gtx.Dp(unit.Dp(20))
					ellipse := clip.Ellipse{Max: image.Pt(size, size)}
					defer ellipse.Push(gtx.Ops).Pop()
					c := q.Theme.Color.Divider
					if q.selected[idx] {
						c = q.Theme.Color.Primary
					}
					paint.FillShape(gtx.Ops, c, ellipse.Op(gtx.Ops))
					return layout.Dimensions{Size: image.Pt(size, size)}
				}),
				layout.Rigid(layout.Spacer{Width: q.Theme.Space.Small}.Layout),
				// Option text.
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(q.Theme.Material, q.Theme.Text.Body, q.Section.Options[idx])
					lbl.Color = optColor
					return lbl.Layout(gtx)
				}),
			)
		})
	})
}

func (q *Quiz) renderResult(gtx layout.Context) layout.Dimensions {
	if q.correct {
		lbl := material.Label(q.Theme.Material, q.Theme.Text.H3, i18n.T("quiz.correct"))
		lbl.Color = q.Theme.Color.Success
		return lbl.Layout(gtx)
	}

	// Wrong answer: show message + retry button.
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(q.Theme.Material, q.Theme.Text.H3, i18n.T("quiz.try_again"))
			lbl.Color = q.Theme.Color.Error
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: q.Theme.Space.Small}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(q.Theme.Material, &q.retryBtn, i18n.T("quiz.retry"))
			btn.Background = q.Theme.Color.Surface
			btn.Color = q.Theme.Color.Text
			return btn.Layout(gtx)
		}),
	)
}

func (q *Quiz) isCorrect() bool {
	for _, idx := range q.Section.Correct {
		if !q.selected[idx] {
			return false
		}
	}
	for i, sel := range q.selected {
		if sel && !q.inCorrect(i) {
			return false
		}
	}
	return true
}

func (q *Quiz) inCorrect(idx int) bool {
	for _, c := range q.Section.Correct {
		if c == idx {
			return true
		}
	}
	return false
}

// SelectOption toggles selection of an option.
func (q *Quiz) SelectOption(idx int) {
	if q.checked || idx < 0 || idx >= len(q.selected) {
		return
	}
	if len(q.Section.Correct) == 1 {
		// Single answer: radio behavior.
		for i := range q.selected {
			q.selected[i] = false
		}
	}
	q.selected[idx] = !q.selected[idx]
}

// IsCorrect reports whether the quiz was answered correctly.
func (q *Quiz) IsCorrect() bool {
	return q.checked && q.correct
}

// ResultColor returns the result feedback color.
func (q *Quiz) ResultColor() color.NRGBA {
	if q.correct {
		return q.Theme.Color.Success
	}
	return q.Theme.Color.Error
}
