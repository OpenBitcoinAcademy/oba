## The Watch-Only Workflow

A hardware wallet signing workflow starts with a watch-only wallet on a networked computer. The watch-only wallet holds the extended public keys (xpubs) for all signers but no private keys. It can generate addresses, track balances, and build transactions, but it cannot sign.

When the user wants to spend, the watch-only wallet creates a PSBT. It fills in the unsigned transaction, attaches WITNESS_UTXO data for each input, and writes BIP32_DERIVATION entries so the hardware wallet knows which keys to use. The PSBT is exported to a file on an SD card or encoded as an animated QR code.

The hardware wallet receives the PSBT, parses the inputs and outputs, and displays a summary to the user: which addresses receive funds, how much goes to each, and how much is paid in fees. The user confirms on the device. The hardware wallet derives the private keys from its seed using the BIP 32 paths in the PSBT, signs each input it can, writes PARTIAL_SIG entries, and exports the updated PSBT back via SD card or QR code.

The watch-only wallet imports the signed PSBT. If all required signatures are present, it finalizes and extracts the raw transaction, then broadcasts it. If more signatures are needed (as in a multisig setup), it passes the PSBT to the next signer.
