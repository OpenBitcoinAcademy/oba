## What a Full Node Does

A full node downloads and independently verifies every block in the blockchain, starting from the genesis block in 2009. It checks every transaction against every consensus rule. It maintains its own copy of the UTXO set (all unspent transaction outputs).

When you run a full node, you do not trust anyone. Your node verifies the entire history of the network by itself. If a block or transaction violates any rule, your node rejects it, regardless of how many other nodes accept it.

A full node also relays valid transactions and blocks to other nodes, contributing to the health and resilience of the network.
