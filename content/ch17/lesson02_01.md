## Two-Round Signing

FROST signing happens in two rounds between the participating signers. A coordinator relays messages but never learns enough to forge a signature or extract secret shares. The coordinator is untrusted.

In round one, each signer generates a pair of random nonces and sends the corresponding nonce commitments to the coordinator. The coordinator collects commitments from all t participating signers and broadcasts them to the group. These commitments bind each signer to their nonces before anyone sees other signers' values, preventing a rogue signer from adaptively choosing nonces to manipulate the outcome.

In round two, each signer computes the aggregate nonce from the collected commitments, constructs the Schnorr challenge hash, and produces a signature share using their secret key share and their nonce. Each signer sends their signature share to the coordinator.
