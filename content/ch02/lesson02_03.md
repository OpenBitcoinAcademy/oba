## From Mempool to Blockchain

Alice's transaction sits in the mempool, waiting for a miner to include it in a block. Miners select transactions from their mempool, prioritizing those with higher fee rates per unit of weight.

About five minutes later, a miner finds a valid proof-of-work solution for a new block that includes Alice's transaction. The miner broadcasts the block. Every full node verifies the block independently: valid proof of work, valid transactions, correct format. Nodes that accept the block add it to their copy of the blockchain.

Alice's transaction now has one confirmation. Bob's wallet updates the status. Each subsequent block built on top of this one adds another confirmation. After six confirmations (about one hour), the transaction is considered highly secure against reversal.
