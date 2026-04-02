## Ark: Off-Chain UTXOs

Ark is a Layer 2 protocol that enables cheap Bitcoin transactions without requiring channels. An Ark Service Provider (ASP) coordinates rounds where users exchange Virtual UTXOs (VTXOs): valid Bitcoin transactions kept off-chain.

Each VTXO is a leaf in a tree of transactions rooted in a single on-chain shared UTXO. Users can always broadcast their VTXO to claim funds on-chain (unilateral exit). During normal operation, the ASP batches thousands of transfers into one on-chain settlement.

Ark does not require opening channels or locking liquidity in advance. Users receive VTXOs by forfeiting old ones in periodic rounds. Atomic swaps via connector outputs ensure no funds are lost during the exchange.
