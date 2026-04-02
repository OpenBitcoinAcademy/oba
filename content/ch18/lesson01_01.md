## Two Transactions for Thousands of Payments

A payment channel is a 2-of-2 multisig contract on the Bitcoin blockchain. Two parties lock funds into a shared output. They exchange signed commitment transactions off-chain to update the balance split. Only two on-chain transactions are needed: one to open the channel and one to close it.

Between open and close, the parties can exchange thousands of payments. Each payment updates the commitment transaction that splits the locked funds. The blockchain never sees these intermediate states. The result: fast payments (milliseconds), near-zero fees, and no block confirmation delay.
