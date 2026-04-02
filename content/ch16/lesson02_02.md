## Updater and Signer

The Updater adds the metadata that Signers need. For each input, the Updater attaches the UTXO being spent (either the full previous transaction or the specific witness UTXO), the BIP 32 derivation path for the signing key, the redeem script if the input is P2SH, and the witness script if the input is P2WSH. For each output, the Updater may attach BIP 32 derivation paths so the signer can verify change outputs belong to the same wallet.

The Signer reads the PSBT, identifies the inputs it can sign, and produces partial signatures. For each input it signs, it writes a PARTIAL_SIG entry keyed by the public key. The Signer does not modify the transaction itself. It does not finalize scripts. It adds its signature and passes the PSBT to the next participant.

A Signer must verify the UTXO data before signing. If the PSBT claims an input spends 0.5 BTC but the actual UTXO holds 1.0 BTC, the signer would unknowingly donate 0.5 BTC to fees. Hardware wallets check UTXO amounts against output amounts and warn the user if the fee seems unreasonable.
