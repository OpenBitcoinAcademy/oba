## The Randomness Requirement

Both ECDSA and Schnorr require a random number (nonce) during signature creation. The security of the entire system depends on this randomness.

If the same nonce is used to sign two different messages with the same private key, the private key can be computed from the two signatures. This is not a theoretical risk. In 2013, a flaw in the Android random number generator caused multiple Bitcoin wallets to reuse nonces, and funds were stolen.

Modern implementations use deterministic nonces (RFC 6979 for ECDSA, BIP340 for Schnorr), deriving the nonce from the private key and message hash. This eliminates the dependency on a random number generator and prevents nonce reuse.
