## Fee Estimation

Wallet software estimates fee rates by analyzing the mempool: the set of unconfirmed transactions waiting for inclusion in a block.

A common approach: look at the fee rates of transactions recently included in blocks and the fee rates of transactions still waiting. Set a fee rate high enough to compete with the lowest-fee-rate transaction in the most recent block, with a margin for safety.

Fee estimation is imperfect. Demand fluctuates. A spike in usage can make a previously adequate fee rate insufficient. Some wallets allow users to set custom fee rates for more control.
