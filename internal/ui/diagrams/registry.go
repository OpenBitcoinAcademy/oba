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
	// Ch01 Introduction.
	"centralized_vs_p2p": &CentralizedVsP2P{},
	// Ch02 How Bitcoin Works.
	"tx_lifecycle": &TxLifecycle{},
	// Ch03 Bitcoin Core.
	"node_architecture": &NodeArchitecture{},
	// Ch04 Keys and Addresses.
	"hash_flow":        &HashFlow{},
	"key_derivation":   &KeyDerivation{},
	"address_pipeline": &AddressPipeline{},
	// Ch05 Wallets.
	"hd_wallet_tree": &HDWalletTree{},
	// Ch06 Transactions.
	"tx_anatomy": &TxAnatomy{},
	"utxo_model": &UTXOModel{},
	// Ch07 Authorization.
	"script_execution": &ScriptExecution{},
	// Ch08 Signatures.
	"signature_flow": &SignatureFlow{},
	// Ch09 Fees.
	"fee_weight": &FeeWeight{},
	// Ch10 Network.
	"network_topology": &NetworkTopology{},
	// Ch11 Blockchain.
	"merkle_tree": &MerkleTree{},
	// Ch12 Mining.
	"mining_process":   &MiningProcess{},
	"halving_schedule": &HalvingSchedule{},
	// Ch14 Taproot.
	"taproot_key_path": &TaprootKeyPath{},
	"taproot_mast":     &TaprootMAST{},
	"security_layers":  &SecurityLayers{},
}

// Validate checks that a diagram ID exists in the registry.
func Validate(id string) error {
	if _, ok := Registry[id]; !ok {
		return fmt.Errorf("diagram %q not found in registry", id)
	}
	return nil
}
