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
	"tx_lifecycle":       &TxLifecycle{},
	"tx_walkthrough":     &TxWalkthrough{},
	"confirmation_depth": &ConfirmationDepth{},
	// Ch03 Bitcoin Core.
	"node_architecture": &NodeArchitecture{},
	// Ch04 Keys and Addresses.
	"hash_flow":        &HashFlow{},
	"key_derivation":   &KeyDerivation{},
	"address_pipeline": &AddressPipeline{},
	// Ch05 Wallets.
	"hd_wallet_tree":   &HDWalletTree{},
	"bip32_derivation": &BIP32Derivation{},
	"seed_recovery":    &SeedRecovery{},
	"key_generation":   &KeyGeneration{},
	// Ch06 Transactions.
	"tx_anatomy":  &TxAnatomy{},
	"utxo_model":  &UTXOModel{},
	"wire_format": &WireFormat{},
	// Ch07 Authorization.
	"script_execution": &ScriptExecution{},
	"multisig_flow":    &MultisigFlow{},
	"segwit_layout":    &SegwitLayout{},
	"taproot_spending": &TaprootSpending{},
	// Ch08 Signatures.
	"signature_flow":   &SignatureFlow{},
	"ecdsa_flow":       &ECDSAFlow{},
	"schnorr_vs_ecdsa": &SchnorrVsECDSA{},
	"sighash_types":    &SighashTypes{},
	// Ch09 Fees.
	"fee_weight":       &FeeWeight{},
	"fee_rate_mempool": &FeeRateMempool{},
	"rbf_bumping":      &RBFBumping{},
	"package_relay":    &PackageRelay{},
	// Ch10 Network.
	"network_topology": &NetworkTopology{},
	"peer_discovery":   &PeerDiscovery{},
	"compact_blocks":   &CompactBlocks{},
	"spv_proof":        &SPVProof{},
	// Ch11 Blockchain.
	"merkle_tree":     &MerkleTree{},
	"block_chain":     &BlockChain{},
	"block_structure": &BlockStructure{},
	// Ch12 Mining.
	"mining_process":   &MiningProcess{},
	"halving_schedule": &HalvingSchedule{},
	"fork_resolution":  &ForkResolution{},
	// Ch14 Taproot.
	"tweaked_key":      &TweakedKey{},
	"taproot_key_path": &TaprootKeyPath{},
	"taproot_mast":     &TaprootMAST{},
	"security_layers":  &SecurityLayers{},
	"tapscript_verify": &TapscriptVerify{},
	// Ch15 Miniscript.
	"miniscript_pipeline": &MiniscriptPipeline{},
	"miniscript_compiler": &MiniscriptCompiler{},
	"output_descriptors":  &OutputDescriptors{},
	"timelocked_recovery": &TimelockedRecovery{},
	// Ch16 PSBTs.
	"psbt_lifecycle":  &PSBTLifecycle{},
	"psbt_roles":      &PSBTRoles{},
	"psbt_maps":       &PSBTMaps{},
	"hardware_wallet": &HardwareWallet{},
	// Ch17 Threshold Signatures.
	"frost_signing": &FrostSigning{},
	"frost_rounds":  &FrostRounds{},
	"chilldkg":      &ChillDKG{},
	"nonce_reuse":   &NonceReuse{},
	// Ch18 Lightning.
	"lightning_channel": &LightningChannel{},
	"htlc_flow":         &HTLCFlow{},
	"lightning_routing": &LightningRouting{},
	"channel_lifecycle": &ChannelLifecycle{},
	// Ch19 Ecosystem.
	"layer_stack":   &LayerStack{},
	"rgb_seals":     &RGBSeals{},
	"fedimint_arch": &FedimintArch{},
	"ark_vtxo":      &ArkVTXO{},
}

// Validate checks that a diagram ID exists in the registry.
func Validate(id string) error {
	if _, ok := Registry[id]; !ok {
		return fmt.Errorf("diagram %q not found in registry", id)
	}
	return nil
}
