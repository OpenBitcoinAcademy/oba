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

// BlockInspector displays a parsed block header with all 6 fields.
type BlockInspector struct {
	Theme  *theme.Theme
	header *bitcoin.BlockHeader
	hash   string
}

// NewBlockInspector creates a block inspector widget.
func NewBlockInspector(th *theme.Theme, cfg *content.ExerciseConfig) *BlockInspector {
	bi := &BlockInspector{Theme: th}
	// Default to genesis block.
	headerHex := "0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c"
	if cfg != nil {
		if h := cfg.ConfigString("block_header_hex", ""); h != "" {
			headerHex = h
		}
	}
	data, err := hex.DecodeString(headerHex)
	if err == nil {
		h, err := bitcoin.ParseBlockHeader(data)
		if err == nil {
			bi.header = h
			bi.hash = h.BlockHash()
		}
	}
	return bi
}

// Layout renders the block inspector.
func (bi *BlockInspector) Layout(gtx layout.Context) layout.Dimensions {
	th := bi.Theme
	return layout.Inset{Top: th.Space.Medium, Bottom: th.Space.Medium}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.block_inspector_title"))
					lbl.Color = th.Color.Text
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if bi.header == nil {
						lbl := material.Label(th.Material, th.Text.Body, "[no block header loaded]")
						lbl.Color = th.Color.Error
						return lbl.Layout(gtx)
					}
					return bi.renderHeader(gtx)
				}),
			)
		})
}

func (bi *BlockInspector) renderHeader(gtx layout.Context) layout.Dimensions {
	th := bi.Theme
	h := bi.header
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(bi.field("Block Hash:", bi.hash)),
		layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
		layout.Rigid(bi.field("Version:", fmt.Sprintf("%d", h.Version))),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
		layout.Rigid(bi.field("Previous Block:", truncHex(h.PrevBlockHex(), 24))),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
		layout.Rigid(bi.field("Merkle Root:", truncHex(h.MerkleRootHex(), 24))),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
		layout.Rigid(bi.field("Timestamp:", h.Time().UTC().Format("2006-01-02 15:04:05 UTC"))),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
		layout.Rigid(bi.field("Bits (Target):", fmt.Sprintf("0x%08x", h.Bits))),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
		layout.Rigid(bi.field("Nonce:", fmt.Sprintf("%d", h.Nonce))),
	)
}

func (bi *BlockInspector) field(label, value string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		th := bi.Theme
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Caption, label)
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						paint.FillShape(gtx.Ops, th.Color.InfoBg,
							clip.Rect{Max: gtx.Constraints.Min}.Op())
						return layout.Dimensions{Size: gtx.Constraints.Min}
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top: unit.Dp(3), Bottom: unit.Dp(3),
							Left: unit.Dp(6), Right: unit.Dp(6),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(th.Material, th.Text.Code, value)
							lbl.Color = th.Color.Text
							lbl.MaxLines = 2
							return lbl.Layout(gtx)
						})
					}),
				)
			}),
		)
	}
}

func truncHex(s string, n int) string {
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}
