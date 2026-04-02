package components

import (
	"encoding/hex"
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/bitcoin"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TxInspector displays the parsed anatomy of a real Bitcoin transaction.
type TxInspector struct {
	Theme *theme.Theme
	tx    *bitcoin.Tx
	txid  string
	size  int
}

// NewTxInspector creates a transaction inspector widget.
func NewTxInspector(th *theme.Theme, cfg *content.ExerciseConfig) *TxInspector {
	ti := &TxInspector{Theme: th}
	if cfg != nil {
		txHex := cfg.ConfigString("preloaded_tx_hex", "")
		if txHex != "" {
			data, err := hex.DecodeString(txHex)
			if err == nil {
				tx, err := bitcoin.DeserializeTx(data)
				if err == nil {
					ti.tx = tx
					ti.txid = tx.TxID()
					ti.size = len(data)
				}
			}
		}
	}
	return ti
}

// Layout renders the transaction inspector.
func (ti *TxInspector) Layout(gtx layout.Context) layout.Dimensions {
	th := ti.Theme

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.tx_inspector_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.tx_inspector_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if ti.tx == nil {
					lbl := material.Label(th.Material, th.Text.Body, "[no transaction loaded]")
					lbl.Color = th.Color.Error
					return lbl.Layout(gtx)
				}
				return ti.renderTx(gtx)
			}),
		)
	})
}

func (ti *TxInspector) renderTx(gtx layout.Context) layout.Dimensions {
	th := ti.Theme
	tx := ti.tx

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// TxID.
		layout.Rigid(ti.field(i18n.T("exercise_ui.tx_inspector_txid"), ti.txid)),
		layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
		// Version + Size.
		layout.Rigid(ti.field(i18n.T("exercise_ui.tx_inspector_version"), fmt.Sprintf("%d", tx.Version))),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
		layout.Rigid(ti.field(
			i18n.TFmt("exercise_ui.tx_inspector_size", map[string]string{"bytes": fmt.Sprintf("%d", ti.size)}),
			"",
		)),
		layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
		// Inputs.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ti.renderInputs(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
		// Outputs.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ti.renderOutputs(gtx)
		}),
	)
}

func (ti *TxInspector) renderInputs(gtx layout.Context) layout.Dimensions {
	th := ti.Theme
	var children []layout.FlexChild
	for idx, in := range ti.tx.Inputs {
		idx := idx
		in := in
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			prevTx := hex.EncodeToString(in.PrevTxID[:])
			if len(prevTx) > 16 {
				prevTx = prevTx[:16] + "..."
			}
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.BodySmall,
						i18n.TFmt("exercise_ui.tx_inspector_input", map[string]string{"index": fmt.Sprintf("%d", idx)}))
					lbl.Color = th.Color.Primary
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(ti.codeField(i18n.T("exercise_ui.tx_inspector_prev_tx")+" "+prevTx)),
				layout.Rigid(ti.codeField(i18n.T("exercise_ui.tx_inspector_prev_index")+" "+fmt.Sprintf("%d", in.PrevIndex))),
				layout.Rigid(ti.codeField(i18n.TFmt("exercise_ui.tx_inspector_script_len", map[string]string{"len": fmt.Sprintf("%d", len(in.ScriptSig))}))),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			)
		}))
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

func (ti *TxInspector) renderOutputs(gtx layout.Context) layout.Dimensions {
	th := ti.Theme
	var children []layout.FlexChild
	for idx, out := range ti.tx.Outputs {
		idx := idx
		out := out
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			sats := fmt.Sprintf("%d", out.Value)
			btc := fmt.Sprintf("%.8f", float64(out.Value)/100000000.0)
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.BodySmall,
						i18n.TFmt("exercise_ui.tx_inspector_output", map[string]string{"index": fmt.Sprintf("%d", idx)}))
					lbl.Color = th.Color.Success
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(ti.codeField(i18n.TFmt("exercise_ui.tx_inspector_value_sat", map[string]string{"sats": sats}))),
				layout.Rigid(ti.codeField(i18n.TFmt("exercise_ui.tx_inspector_value_btc", map[string]string{"btc": btc}))),
				layout.Rigid(ti.codeField(i18n.TFmt("exercise_ui.tx_inspector_script_len", map[string]string{"len": fmt.Sprintf("%d", len(out.ScriptPubKey))}))),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			)
		}))
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

func (ti *TxInspector) field(label, value string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		if value == "" {
			lbl := material.Label(ti.Theme.Material, ti.Theme.Text.BodySmall, label)
			lbl.Color = ti.Theme.Color.TextMuted
			return lbl.Layout(gtx)
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(ti.Theme.Material, ti.Theme.Text.Caption, label)
				lbl.Color = ti.Theme.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						paint.FillShape(gtx.Ops, ti.Theme.Color.InfoBg,
							clip.Rect{Max: gtx.Constraints.Min}.Op())
						return layout.Dimensions{Size: gtx.Constraints.Min}
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top: unit.Dp(4), Bottom: unit.Dp(4),
							Left: unit.Dp(8), Right: unit.Dp(8),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(ti.Theme.Material, ti.Theme.Text.Code, value)
							lbl.Color = ti.Theme.Color.Text
							lbl.MaxLines = 2
							return lbl.Layout(gtx)
						})
					}),
				)
			}),
		)
	}
}

func (ti *TxInspector) codeField(text string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		lbl := material.Label(ti.Theme.Material, ti.Theme.Text.Code, text)
		lbl.Color = ti.Theme.Color.Text
		return layout.Inset{Left: unit.Dp(8), Top: unit.Dp(2)}.Layout(gtx, lbl.Layout)
	}
}
