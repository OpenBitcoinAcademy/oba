## Compact Block Relay

Defined in **BIP 152**, compact block relay reduces the bandwidth needed to propagate a new block. The key insight: most transactions in a new block are already in the receiving node's mempool.

Instead of sending the full block, the announcing node sends a `cmpctblock` message. This contains the block header, a short transaction ID list, and the coinbase transaction. Short transaction IDs are 6-byte truncated hashes that let the receiver match transactions it already has.

The receiving node reconstructs the block from its own mempool using the short IDs. If any transactions are missing, it requests them with a `getblocktxn` message and receives them in a `blocktxn` response.

BIP 152 defines two modes. In **low-bandwidth mode**, a node first announces the block with an `inv` message. The peer requests the compact block only if interested. In **high-bandwidth mode**, the node sends the compact block immediately without waiting for a request. Miners and well-connected nodes typically use high-bandwidth mode to minimize latency.

Compact block relay cuts propagation bandwidth by 90% or more for a typical block. Faster propagation gives smaller miners a fairer chance, strengthening decentralization.
