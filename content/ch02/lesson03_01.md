## The Mining Lottery

Mining is a decentralized lottery. Each miner creates a candidate block from transactions in their mempool. They hash the block header repeatedly, changing a number called the nonce each time, looking for a hash value below a target set by the network.

Finding a valid hash requires trillions of attempts. Verifying a valid hash requires computing it once. This asymmetry is the core of proof of work: creating a valid block is expensive, checking one is cheap.

The miner who finds a valid hash first broadcasts the block. Other miners verify it, accept it, and immediately begin working on the next block. The winner collects the block reward: newly created bitcoins (the subsidy) plus the sum of all transaction fees in the block.
