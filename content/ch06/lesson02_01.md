## Selecting UTXOs

Building a transaction starts with choosing which outputs to spend. Your wallet scans the blockchain for UTXOs locked to your addresses. These are the coins you control.

To send 0.5 BTC, your wallet selects one or more UTXOs that sum to at least 0.5 BTC. If the smallest UTXO is 0.8 BTC, that single UTXO becomes the input. The excess (minus fees) returns to you as change.

If no single UTXO is large enough, the wallet combines multiple UTXOs as separate inputs in the same transaction. Each input requires its own signature.
