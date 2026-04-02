## Transaction Fees

Bitcoin transactions do not have an explicit fee field. The fee is implicit: it is the difference between the total value of all inputs and the total value of all outputs.

If your inputs total 100,000 satoshis and your outputs total 99,800 satoshis, the fee is 200 satoshis. Miners collect this fee as an incentive to include your transaction in a block.

Higher fees mean faster confirmation. When the network is busy, miners prioritize transactions with higher fee rates. A transaction that pays too little may wait hours or days.

The fee rate depends on the transaction weight (measured in virtual bytes, or vbytes), not the amount being sent. Segwit introduced weight units where witness data (signatures) is discounted compared to other transaction data. A transaction sending 0.001 BTC can cost the same fee as one sending 1,000 BTC, if both have the same structure.
