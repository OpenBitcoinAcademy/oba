## Syncing the Blockchain

A new node must download and validate the entire blockchain before it can verify current transactions. This process is called **Initial Block Download** (IBD).

The node sends `getheaders` messages to its peers, requesting block headers in batches. Headers are small (80 bytes each) and arrive quickly. The node uses them to build the chain of headers with the most cumulative proof-of-work.

Once the node identifies the best header chain, it requests full blocks using `getdata` messages. It downloads blocks from multiple peers in parallel to speed up the process. Each block is validated against the consensus rules as it arrives: transaction signatures, script execution, weight limits, and the proof-of-work target.

IBD can take hours or days on slow hardware. After it completes, the node switches to normal operation. It receives new blocks and transactions as they are broadcast and validates them in real time.
