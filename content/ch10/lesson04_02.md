## Bloom Filters and Their Privacy Problem

**BIP 37** introduced Bloom filters as a way for SPV clients to request only relevant transactions from full nodes. A Bloom filter is a probabilistic data structure. The client creates a filter containing its addresses and sends it to a peer. The peer sends only transactions matching the filter.

The design has a serious privacy flaw. The full node sees the Bloom filter and can deduce which addresses belong to the client. Even though a Bloom filter has false positives, research has shown that a malicious node can identify the client's addresses with high accuracy by analyzing the filter's bit pattern.

A node serving Bloom filter requests also bears a heavy computational cost. It must test every transaction in every block against the filter. This creates a denial-of-service vector: an attacker can connect many times with different filters and force the node to do expensive work. Many node operators now disable BIP 37 support.
