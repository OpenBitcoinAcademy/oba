## Derivation Paths (BIP44)

BIP44 defines a standard structure for the key tree. The path format is: m / purpose' / coin_type' / account' / change / address_index.

**purpose** is 44 for standard addresses (or 84 for SegWit, 86 for Taproot). **coin_type** is 0 for Bitcoin mainnet, 1 for testnet. **account** allows separating funds into logical groups. **change** is 0 for receiving addresses and 1 for change addresses. **address_index** increments for each new address.

A typical Bitcoin receiving address path: m/84'/0'/0'/0/0. This means: SegWit purpose, Bitcoin mainnet, first account, receiving (not change), first address.

Standardized paths let different wallet software reconstruct the same set of keys from the same seed.
