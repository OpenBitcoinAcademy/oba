## Conditional Payments Across Hops

Hash Time-Locked Contracts (HTLCs) enable payments between parties that do not share a direct channel. An HTLC combines two conditions: a hashlock (reveal a secret to claim funds) and a timelock (if the secret is not revealed in time, the funds return to the sender).

Carol generates a random secret R and sends the hash H(R) to Alice in a payment invoice. Alice does not know R. She creates an HTLC with Bob: "Pay 1.001 BTC if you reveal the preimage of H(R) within 40 blocks." Bob creates an HTLC with Carol: "Pay 1.000 BTC if you reveal the preimage of H(R) within 30 blocks."
