## Onion Routing

Lightning wraps payment instructions in layers of encryption, like an onion. Each forwarding node decrypts one layer, which reveals only the next hop and the amount to forward. The node does not learn the sender, the final recipient, or the total path length.

This privacy model (based on SPHINX) means that Bob, forwarding a payment from Alice to Carol, knows only that Alice sent him something and that he should forward it to Carol. He does not know if Alice is the original sender or another forwarding node. He does not know if Carol is the final recipient.

The onion has a fixed size regardless of the number of hops, preventing path-length analysis.
