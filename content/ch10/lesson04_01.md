## SPV Clients

Not every device can store and validate the full blockchain. Mobile phones, embedded devices, and hardware wallets have limited storage and processing power. These devices run **Simplified Payment Verification** (SPV) clients.

An SPV client downloads only block headers, not full blocks. The header chain is small: about 60 MB for the entire history of Bitcoin. The client can verify that a block header meets the proof-of-work target, confirming that a miner spent real energy producing it.

To check if a transaction is confirmed, the SPV client requests a **Merkle proof** from a full node. The proof shows that the transaction's hash is included in a block's Merkle root. The client trusts that full nodes validated the block's transactions, because faking a valid proof-of-work header is prohibitively expensive.

SPV gives lightweight clients a reasonable level of assurance without downloading the full blockchain. The tradeoff is that an SPV client cannot detect if a block contains an invalid transaction. It trusts the economic weight of proof-of-work rather than verifying every rule itself.
