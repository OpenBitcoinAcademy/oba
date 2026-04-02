## Taproot (P2TR)

Taproot (BIP341) is the newest output type, activated in November 2021 as segwit version 1. A Taproot output commits to a single public key. Spending can happen in two ways.

**Key path**: the owner signs directly with the committed key. On the blockchain, this looks identical to a single-signature spend. No script is revealed. No one can tell whether the output had additional spending conditions.

**Script path**: the owner reveals a script from a Merkle tree of scripts committed to in the output. This allows complex conditions (multisig, timelocks, hash locks) while keeping the common case (key path) compact and private.

Taproot addresses start with "bc1p" on mainnet and use bech32m encoding.
