## Package Relay

Historically, Bitcoin Core evaluated each transaction independently for mempool admission. A low-fee parent transaction would be rejected even if a high-fee child transaction depended on it.

Package relay changes this. Nodes evaluate groups of related transactions together, accepting a low-fee parent if its child brings the combined fee rate above the threshold.

This improves CPFP reliability. Without package relay, the parent might not propagate to miners, making the child useless. With package relay, the parent and child travel together through the network.
