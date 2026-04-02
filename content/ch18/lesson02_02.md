## Preimage Propagation

Carol knows R. She reveals it to Bob and claims 1.000 BTC. Bob now knows R. He reveals it to Alice and claims 1.001 BTC. The preimage propagates backward along the payment path.

Alice paid 1.001 BTC. Carol received 1.000 BTC. Bob earned 0.001 BTC as a routing fee. The payment settled in seconds without touching the blockchain.

All HTLCs in the chain use the same hash H(R). Either the preimage propagates all the way back (the payment succeeds entirely) or the timelocks expire (the payment fails entirely, all funds returned). No intermediate state is possible. The payment is atomic.
