## How Blocks Propagate

When a miner finds a valid block, it broadcasts the block to its peers. Those peers validate it and forward it to their peers. The block spreads across the network in a few seconds.

Speed matters. If two miners find valid blocks at nearly the same time, the block that reaches more of the network first has an advantage. Miners who hear about block A first will build on top of it. Miners who hear about block B first will build on that. The race resolves when the next block extends one chain, and the other block becomes stale.

Slow propagation hurts decentralization. If large mining pools with fast connections hear about blocks first, they have a head start on the next block. Smaller miners with slower connections waste work on stale blocks more often. The network protocol addresses this asymmetry through compact block relay.
