## Multisig Passing and Transport

In a multisig configuration, the PSBT travels through each signer in sequence or is distributed to all signers in parallel. Consider a 2-of-3 multisig: the coordinator creates the PSBT and sends it to signer A and signer B simultaneously. Each signer adds its partial signature and returns the PSBT. The coordinator combines both PSBTs, finalizes, and broadcasts.

The transport mechanism is not specified by BIP 174. PSBTs can move on SD cards, as QR codes (using the UR encoding from BCR-2020-005), over NFC, via email, through a file sharing service, or by any other channel. The format is self-contained. Each signer gets everything it needs from the PSBT itself, with no side-channel required.

PSBT is plaintext, not encrypted. Anyone who intercepts a PSBT can see the transaction amounts, the addresses involved, and the partial signatures collected so far. This is a privacy concern, not a security concern. An attacker who reads a PSBT cannot steal funds because the PSBT does not contain private keys. But the attacker learns what the user is spending and where the funds go. For sensitive transactions, the PSBT should be transported over an encrypted channel or on physical media that the user controls.
