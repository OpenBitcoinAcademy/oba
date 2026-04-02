## Finding Peers

When a Bitcoin node starts for the first time, it knows no other nodes. It needs to find at least one peer to begin participating in the network.

Bitcoin uses **DNS seeds** for initial discovery. DNS seeds are DNS servers operated by Bitcoin community members. They return the IP addresses of known, stable full nodes. Bitcoin Core has several DNS seeds hardcoded in its source code.

If DNS is unavailable (blocked or censored), Bitcoin Core also includes a list of hardcoded IP addresses as a last resort. These addresses are updated with each software release.

Once a node connects to its first peer, it can ask that peer for additional addresses. The node builds a diverse set of connections, typically 8 outbound and up to 125 inbound.
