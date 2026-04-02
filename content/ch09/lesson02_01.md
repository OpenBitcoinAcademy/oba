## Fee Rates

Miners compare transactions by their fee rate: the fee divided by the transaction weight. The unit is satoshis per virtual byte (sat/vB).

A transaction paying 1,000 satoshis with a weight of 500 vbytes has a fee rate of 2 sat/vB. A transaction paying 2,000 satoshis with a weight of 1,000 vbytes also has a fee rate of 2 sat/vB. Miners treat them equally.

During quiet periods, 1 sat/vB may be enough for next-block confirmation. During high demand, fee rates can exceed 100 sat/vB. Wallet software estimates the appropriate fee rate based on current mempool conditions.
