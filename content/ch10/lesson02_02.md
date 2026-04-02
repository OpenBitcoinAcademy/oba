## The Handshake

Two nodes establish a connection with a version/verack handshake over TCP port 8333 (the default Bitcoin network port).

The connecting node sends a `version` message containing its protocol version, block height, current time, and the services it supports. The receiving node replies with its own `version` message. Each node then sends a `verack` (version acknowledgement) to confirm.

After the handshake completes, the nodes can exchange data. If the protocol versions are incompatible, the connection is dropped.

Nodes share known peer addresses using `addr` and `getaddr` messages. When a node learns new addresses, it stores them and may connect to them later. This gossip protocol lets the network grow and heal without any central directory.
