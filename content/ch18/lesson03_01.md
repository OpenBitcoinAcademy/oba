## Source-Based Routing

The sender, not the network, chooses the payment path. Each Lightning node maintains a local view of the network topology (which nodes exist, which channels connect them, their capacities and fee rates). The sender computes a route from this information.

This differs from internet routing, where each router independently decides the next hop. In Lightning, the sender has full control over the path. The intermediate nodes follow forwarding instructions without knowing the full route.
