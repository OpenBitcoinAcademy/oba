## Opening a Channel

Opening a channel requires one on-chain transaction: the funding transaction. It creates a 2-of-2 multisig output. Before broadcasting, both parties sign the first commitment transaction (the safety net that returns funds if the channel fails).

The channel is open once the funding transaction is confirmed. Both parties can begin updating the balance by exchanging new commitment transactions. The on-chain cost is one transaction fee for the funding transaction.
