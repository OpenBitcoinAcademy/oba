## Native SegWit (P2WPKH and P2WSH)

Segregated Witness moved signature data out of the input script and into a separate witness structure. This fixed transaction malleability (third parties could modify the txid by tweaking the signature encoding) and enabled the witness discount for fees.

P2WPKH (Pay to Witness Public Key Hash) is the segwit equivalent of P2PKH. The output script contains: OP_0 <20-byte-pubkey-hash>. The witness provides the signature and public key. The input script is empty.

P2WSH (Pay to Witness Script Hash) is the segwit equivalent of P2SH. The output script contains: OP_0 <32-byte-script-hash>. The witness provides the script and its required data. P2WSH uses a 32-byte SHA-256 hash instead of P2SH's 20-byte HASH160, providing stronger collision resistance.
