## Propagation Across the Network

Alice's wallet signs the transaction with her private key and broadcasts it to the Bitcoin network. Her wallet connects to several nodes, which verify the transaction independently: correct format, valid signature, unspent inputs, output values do not exceed input values.

Each node that accepts the transaction adds it to its mempool (a local list of unconfirmed transactions) and relays it to its peers. Within seconds, the transaction reaches thousands of nodes across the network. No central server coordinates this. Each node follows the same rules independently.

Bob's wallet monitors the network for transactions paying to his address. When it sees Alice's transaction, it shows as "unconfirmed" in Bob's wallet. The payment has been broadcast but not yet recorded in a block.
