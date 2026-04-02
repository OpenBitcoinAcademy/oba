package components

import (
	"encoding/hex"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"

	"github.com/openbitcoinacademy/oba/internal/bitcoin"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SigVerifier generates a key pair, signs a message, and lets users
// verify the signature against the public key.
type SigVerifier struct {
	Theme     *theme.Theme
	signBtn   widget.Clickable
	verifyBtn widget.Clickable
	resetBtn  widget.Clickable

	message   string
	pubKeyHex string
	sigHex    string
	signed    bool
	verified  *bool // nil=not yet, true=valid, false=invalid

	// Raw bytes for verification.
	msgHash [32]byte
	sigDER  []byte
	pubKey  []byte
}

// NewSigVerifier creates a signature verifier widget.
func NewSigVerifier(th *theme.Theme, cfg *content.ExerciseConfig) *SigVerifier {
	msg := "Hello, Bitcoin"
	if cfg != nil {
		msg = cfg.ConfigString("default_message", msg)
	}
	return &SigVerifier{
		Theme:   th,
		message: msg,
	}
}

// Layout renders the signature verifier.
func (sv *SigVerifier) Layout(gtx layout.Context) layout.Dimensions {
	th := sv.Theme

	if sv.signBtn.Clicked(gtx) {
		sv.doSign()
	}
	if sv.verifyBtn.Clicked(gtx) && sv.signed {
		sv.doVerify()
	}
	if sv.resetBtn.Clicked(gtx) {
		sv.signed = false
		sv.verified = nil
		sv.pubKeyHex = ""
		sv.sigHex = ""
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.sig_verify_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Prompt.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.sig_verify_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Sign button or results.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !sv.signed {
					btn := material.Button(th.Material, &sv.signBtn, i18n.T("exercise.generate"))
					btn.Background = th.Color.Primary
					btn.Color = th.Color.OnPrimary
					return btn.Layout(gtx)
				}
				return sv.renderResults(gtx)
			}),
		)
	})
}

func (sv *SigVerifier) renderResults(gtx layout.Context) layout.Dimensions {
	th := sv.Theme
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Message.
		layout.Rigid(sv.fieldDisplay("Message", sv.message, th.Color.InfoBg)),
		layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
		// Public key.
		layout.Rigid(sv.fieldDisplay("Public Key", sv.pubKeyHex, th.Color.TipBg)),
		layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
		// Signature.
		layout.Rigid(sv.fieldDisplay("Signature (DER)", sv.sigHex, th.Color.WarningBg)),
		layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

		// Verify button and result.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if sv.verified == nil {
						btn := material.Button(th.Material, &sv.verifyBtn, i18n.T("exercise_ui.sig_verify_title"))
						btn.Background = th.Color.Accent
						btn.Color = th.Color.OnPrimary
						return btn.Layout(gtx)
					}
					return layout.Dimensions{}
				}),
				layout.Rigid(layout.Spacer{Width: th.Space.Small}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if sv.verified == nil {
						return layout.Dimensions{}
					}
					var text string
					var c = th.Color.Error
					if *sv.verified {
						text = i18n.T("exercise_ui.sig_verify_valid")
						c = th.Color.Success
					} else {
						text = i18n.T("exercise_ui.sig_verify_invalid")
					}
					lbl := material.Label(th.Material, th.Text.H3, text)
					lbl.Color = c
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if sv.verified != nil {
						btn := material.Button(th.Material, &sv.resetBtn, i18n.T("exercise.reset"))
						btn.Background = th.Color.Surface
						btn.Color = th.Color.Primary
						return btn.Layout(gtx)
					}
					return layout.Dimensions{}
				}),
			)
		}),
		layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise.safety_note"))
			lbl.Color = th.Color.TextMuted
			return lbl.Layout(gtx)
		}),
	)
}

func (sv *SigVerifier) fieldDisplay(label, value string, bg interface{}) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		th := sv.Theme
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Caption, label)
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
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
							lbl := material.Label(th.Material, th.Text.Code, value)
							lbl.Color = th.Color.Text
							lbl.MaxLines = 3
							return lbl.Layout(gtx)
						})
					}),
				)
			}),
		)
	}
}

func (sv *SigVerifier) doSign() {
	key, err := bitcoin.GenerateKey()
	if err != nil {
		return
	}

	sv.pubKey = key.PublicKey().SerializeCompressed()
	sv.pubKeyHex = hex.EncodeToString(sv.pubKey)

	// Hash the message (ECDSA signs a hash, not raw data).
	sv.msgHash = bitcoin.SHA256([]byte(sv.message))

	// Reconstruct btcec private key from raw bytes to access Sign.
	privKey, _ := btcec.PrivKeyFromBytes(key.Bytes())
	sig := ecdsa.Sign(privKey, sv.msgHash[:])

	sv.sigDER = sig.Serialize()
	sv.sigHex = hex.EncodeToString(sv.sigDER)
	sv.signed = true
	sv.verified = nil
}

func (sv *SigVerifier) doVerify() {
	sig, err := ecdsa.ParseDERSignature(sv.sigDER)
	if err != nil {
		f := false
		sv.verified = &f
		return
	}

	pub, err := btcec.ParsePubKey(sv.pubKey)
	if err != nil {
		f := false
		sv.verified = &f
		return
	}

	valid := sig.Verify(sv.msgHash[:], pub)
	sv.verified = &valid
}
