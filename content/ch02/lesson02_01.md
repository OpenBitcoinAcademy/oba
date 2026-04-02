## Alice Pays Bob

Alice wants to buy a product from Bob's online store. Bob's checkout page displays a Bitcoin payment option with an amount and a Bitcoin address.

Alice opens her wallet app. It scans the blockchain for unspent transaction outputs (UTXOs) locked to her addresses. Her wallet finds a UTXO worth 0.015 BTC, more than enough to cover Bob's price of 0.01 BTC.

Alice's wallet constructs a transaction with one input (referencing her 0.015 BTC UTXO) and two outputs: 0.01 BTC to Bob's address, and 0.0049 BTC back to Alice as change. The remaining 0.0001 BTC is the transaction fee.
