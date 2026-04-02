package components

import (
	"encoding/hex"
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/bitcoin"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TxBuilder walks users through constructing a transaction step by step.
type TxBuilder struct {
	Theme    *theme.Theme
	nextBtn  widget.Clickable
	resetBtn widget.Clickable
	step     int
	steps    []txBuildStep
	utxoVal  int64
}

type txBuildStep struct {
	label  string
	output string
}

// NewTxBuilder creates a transaction builder widget.
func NewTxBuilder(th *theme.Theme, cfg *content.ExerciseConfig) *TxBuilder {
	utxoVal := int64(100000) // default: 100,000 satoshis
	if cfg != nil {
		utxoVal = cfg.ConfigInt("fake_utxo_value", 100000)
	}
	return &TxBuilder{Theme: th, utxoVal: utxoVal}
}

// Layout renders the transaction builder.
func (tb *TxBuilder) Layout(gtx layout.Context) layout.Dimensions {
	th := tb.Theme

	if tb.nextBtn.Clicked(gtx) {
		tb.advance()
	}
	if tb.resetBtn.Clicked(gtx) {
		tb.step = 0
		tb.steps = nil
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.tx_builder_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if tb.step == 0 {
					lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.tx_builder_prompt"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}
				return tb.renderSteps(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if tb.step >= 7 {
					btn := material.Button(th.Material, &tb.resetBtn, i18n.T("exercise.reset"))
					btn.Background = th.Color.Surface
					btn.Color = th.Color.Primary
					return btn.Layout(gtx)
				}
				return layout.Flex{}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &tb.nextBtn, i18n.T("exercise.next_step"))
						btn.Background = th.Color.Primary
						btn.Color = th.Color.OnPrimary
						return btn.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Small}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if tb.step > 0 {
							lbl := material.Label(th.Material, th.Text.Caption,
								i18n.TFmt("exercise_ui.tx_builder_step", map[string]string{
									"current": fmt.Sprintf("%d", tb.step),
									"total":   "7",
								}))
							lbl.Color = th.Color.TextMuted
							return lbl.Layout(gtx)
						}
						return layout.Dimensions{}
					}),
				)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if tb.step > 0 {
					lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise.safety_note"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}
				return layout.Dimensions{}
			}),
		)
	})
}

func (tb *TxBuilder) advance() {
	tb.step++
	if tb.step == 1 {
		tb.computeAll()
	}
}

func (tb *TxBuilder) computeAll() {
	// Step 1: Generate sender key pair.
	senderKey, err := bitcoin.GenerateKey()
	if err != nil {
		return
	}
	senderPub := senderKey.PublicKey()
	tb.steps = append(tb.steps, txBuildStep{
		label:  "1. " + i18n.T("exercise_ui.tx_builder_utxo"),
		output: i18n.TFmt("exercise_ui.tx_builder_utxo", map[string]string{"sats": fmt.Sprintf("%d", tb.utxoVal)}),
	})

	// Step 2: Generate recipient.
	recipientKey, err := bitcoin.GenerateKey()
	if err != nil {
		return
	}
	recipientAddr := bitcoin.P2PKH(recipientKey.PublicKey(), false)
	tb.steps = append(tb.steps, txBuildStep{
		label:  "2. " + i18n.T("exercise_ui.tx_builder_recipient"),
		output: recipientAddr,
	})

	// Step 3: Calculate fee.
	fee := int64(226) // typical P2PKH tx size in bytes * 1 sat/byte
	tb.steps = append(tb.steps, txBuildStep{
		label:  "3. " + i18n.T("exercise_ui.tx_builder_fee"),
		output: i18n.TFmt("exercise_ui.tx_builder_fee", map[string]string{"sats": fmt.Sprintf("%d", fee)}),
	})

	// Step 4: Calculate change.
	sendAmount := tb.utxoVal / 2
	change := tb.utxoVal - sendAmount - fee
	tb.steps = append(tb.steps, txBuildStep{
		label:  "4. " + i18n.T("exercise_ui.tx_builder_change"),
		output: i18n.TFmt("exercise_ui.tx_builder_change", map[string]string{"sats": fmt.Sprintf("%d", change)}),
	})

	// Step 5: Build transaction.
	prevHash := bitcoin.SHA256(senderPub.SerializeCompressed())
	fakePrevTx := prevHash

	recipientPubHash := recipientKey.PublicKey().Hash160()
	senderPubHash := senderPub.Hash160()

	tx := &bitcoin.Tx{
		Version: 1,
		Inputs: []bitcoin.TxIn{
			{
				PrevTxID:  fakePrevTx,
				PrevIndex: 0,
				ScriptSig: bitcoin.BuildP2PKHScriptSig([]byte{0x30}, senderPub.SerializeCompressed()),
				Sequence:  0xffffffff,
			},
		},
		Outputs: []bitcoin.TxOut{
			{Value: uint64(sendAmount), ScriptPubKey: bitcoin.BuildP2PKHScript(recipientPubHash)},
			{Value: uint64(change), ScriptPubKey: bitcoin.BuildP2PKHScript(senderPubHash)},
		},
		Locktime: 0,
	}

	tb.steps = append(tb.steps, txBuildStep{
		label:  "5. " + i18n.T("exercise_ui.tx_builder_serialized"),
		output: tx.SerializedHex(),
	})

	// Step 6: Size.
	tb.steps = append(tb.steps, txBuildStep{
		label:  "6. " + i18n.TFmt("exercise_ui.tx_inspector_size", map[string]string{"bytes": fmt.Sprintf("%d", tx.Size())}),
		output: "",
	})

	// Step 7: TxID.
	tb.steps = append(tb.steps, txBuildStep{
		label:  "7. " + i18n.T("exercise_ui.tx_builder_txid"),
		output: tx.TxID(),
	})

	_ = hex.EncodeToString // keep import
}

func (tb *TxBuilder) renderSteps(gtx layout.Context) layout.Dimensions {
	th := tb.Theme
	var children []layout.FlexChild

	visible := tb.step
	if visible > len(tb.steps) {
		visible = len(tb.steps)
	}

	for i := 0; i < visible; i++ {
		s := tb.steps[i]
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.BodySmall, s.label)
					lbl.Color = th.Color.Primary
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if s.output == "" {
						return layout.Dimensions{}
					}
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							paint.FillShape(gtx.Ops, th.Color.InfoBg,
								clip.Rect{Max: gtx.Constraints.Min}.Op())
							return layout.Dimensions{Size: gtx.Constraints.Min}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Top: unit.Dp(4), Bottom: unit.Dp(4),
								Left: unit.Dp(8), Right: unit.Dp(8),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.Code, s.output)
								lbl.Color = th.Color.Text
								lbl.MaxLines = 3
								return lbl.Layout(gtx)
							})
						}),
					)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			)
		}))
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}
