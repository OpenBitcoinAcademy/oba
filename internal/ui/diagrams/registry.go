// Package diagrams maps diagram string IDs (from TOML) to Gio-native
// diagram implementations. Unknown IDs cause startup failure.
package diagrams

import (
	"fmt"

	"gioui.org/layout"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// DiagramWidget is implemented by all Gio-native diagrams.
type DiagramWidget interface {
	Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions
}

// Registry maps diagram IDs to their implementations.
var Registry = map[string]DiagramWidget{
	// Chapter 1.
	"centralized_vs_p2p": &CentralizedVsP2P{},
	"hash_flow":          &HashFlow{},
	"key_derivation":     &KeyDerivation{},
	"address_pipeline":   &AddressPipeline{},
	// Chapter 2.
	"tx_anatomy":       &TxAnatomy{},
	"script_execution": &ScriptExecution{},
	"utxo_model":       &UTXOModel{},
}

// Validate checks that a diagram ID exists in the registry.
func Validate(id string) error {
	if _, ok := Registry[id]; !ok {
		return fmt.Errorf("diagram %q not found in registry", id)
	}
	return nil
}
