## Conflicting Transactions

Alice can create two transactions that spend the same output: one paying Bob and one paying Carol. Both are valid signatures, but only one can be included in the blockchain (the output can be spent only once).

Miners decide which conflicting transaction to include. By default, most miners follow a "first seen" policy and include the transaction they received first. But this is a policy, not a consensus rule. A miner can include any valid transaction.

This is why Bob should wait for confirmations before considering a payment final. The fee Alice pays incentivizes miners to include her transaction promptly, reducing the window for conflicting transactions.
