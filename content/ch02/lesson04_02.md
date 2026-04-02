## The Mempool

When you broadcast a transaction, it does not go into a block immediately. It enters the mempool: a waiting area where unconfirmed transactions sit until a miner includes them in a block.

Every node maintains its own mempool. Transactions propagate across the network in seconds, but confirmation (inclusion in a block) takes an average of 10 minutes. During busy periods, the mempool grows and transactions with low fees may wait longer.

A transaction in the mempool is unconfirmed. It has been validated by nodes (correct signatures, unspent inputs) but has not been written into the blockchain yet.
