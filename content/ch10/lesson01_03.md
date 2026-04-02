## Archival and Pruned Nodes

A full node that keeps every block since the genesis block is called an **archival node**. It can serve historical block data to any peer that requests it. New nodes joining the network rely on archival nodes to download the full blockchain.

Running an archival node requires hundreds of gigabytes of disk space. As of 2024, the blockchain exceeds 600 GB. Not every full node operator can afford that storage.

A **pruned node** validates every block but discards old block data after processing it. It keeps only the UTXO set and the most recent blocks. A pruned node enforces every consensus rule. It does not rely on trust. The tradeoff: it cannot serve historical blocks to other peers.

Both archival and pruned nodes are full nodes. Both validate everything. The difference is whether they store old blocks for others to download.
