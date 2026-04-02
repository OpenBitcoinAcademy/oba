## Node Types

Not every node runs the same software or performs the same role. Nodes differ in what data they store and what functions they provide.

A **full node** downloads and validates every block and every transaction against the consensus rules. It trusts no one. It can independently verify any payment. Bitcoin Core is the most common full node implementation.

A **mining node** is a full node that also competes to create new blocks. It assembles candidate blocks from the mempool and performs proof-of-work. Mining nodes earn the block subsidy and transaction fees when they find a valid block.

A **lightweight node** (also called an SPV client) does not download full blocks. It downloads only block headers and requests proof that specific transactions exist. Lightweight nodes trust full nodes to some degree, trading security for lower resource use.
