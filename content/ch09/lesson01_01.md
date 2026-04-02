## Who Pays the Transaction Fee?

The sender pays the transaction fee. The fee is not a separate payment. It is the difference between the total value of inputs and the total value of outputs.

If Alice creates a transaction with inputs totaling 100,000 satoshis and outputs totaling 99,500 satoshis, the remaining 500 satoshis is the fee. The miner who includes the transaction in a block collects it.

There is no rule requiring a minimum fee. But miners prioritize transactions with higher fee rates per unit of weight. A transaction with too low a fee may wait in the mempool for hours, days, or never confirm.
