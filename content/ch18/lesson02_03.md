## Timelock Decrements

Each hop in the chain has a shorter timelock than the previous hop. Alice gives Bob 40 blocks. Bob gives Carol 30 blocks. This decrement (called the CLTV delta) ensures that each forwarding node has time to claim funds from its downstream hop before its upstream HTLC expires.

If Carol does not reveal R within 30 blocks, Bob's HTLC expires and funds return to Bob. Bob still has 10 blocks to settle with Alice. If Bob fails to reveal R to Alice within 40 blocks, Alice's funds return to her.

The decrement prevents a timing attack where a downstream node delays revealing the preimage until the upstream HTLC expires, trapping the forwarding node.
