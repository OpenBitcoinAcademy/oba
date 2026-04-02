## Compact Block Filters

**BIP 157** and **BIP 158** define a better approach called Compact Block Filters. Instead of the client sending its filter to the server, the server builds a filter for each block and sends it to the client.

Each filter is a compact representation of all the addresses and scripts in a block. The client downloads the filter and checks locally whether any of its addresses appear. If a match is found, the client downloads the full block to extract the relevant transactions.

The privacy advantage is that the server never learns which addresses the client cares about. The server sends the same filter to every client. The client does all the matching locally.

The bandwidth cost is higher than Bloom filters because the client downloads a filter for every block. But the filters are small (about 20 KB per block on average) and can be verified against a hash commitment in the block header chain. The client does not need to trust the server to provide accurate filters.

Running Bitcoin over **Tor** adds another layer of privacy. Tor hides the client's IP address from the nodes it connects to. Combined with Compact Block Filters, a lightweight client can check its balance without revealing its identity or its addresses to any peer.
