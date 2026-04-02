## Fee Sniping

Fee sniping is a theoretical attack where a miner, instead of building on the latest block, re-mines a previous block to steal its transaction fees. If the fees in block N are high enough, a miner might find it more profitable to fork the chain at block N-1 and collect those fees themselves.

Bitcoin Core defends against this by setting the locktime of new transactions to the current block height. A transaction locked to block 800,000 cannot be included in block 799,999. This makes it less profitable to re-mine old blocks, because the new transactions created since then would be unavailable.

Fee sniping has not occurred on the Bitcoin network. The defense is precautionary.
