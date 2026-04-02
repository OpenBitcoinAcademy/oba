## Locking Scripts

Every transaction output contains a locking script, called a scriptPubKey. This script defines the conditions that must be met to spend the output.

The simplest locking script to understand is Pay-to-Public-Key-Hash (P2PKH). It says: "To spend this output, prove you own the private key corresponding to this public key hash." Modern transactions use newer formats (P2WPKH for SegWit, P2TR for Taproot), but P2PKH demonstrates the core concept most clearly.

In script notation, a P2PKH lock looks like: OP_DUP OP_HASH160 <pubkey_hash> OP_EQUALVERIFY OP_CHECKSIG. Each piece is an instruction that the Bitcoin script engine executes.
