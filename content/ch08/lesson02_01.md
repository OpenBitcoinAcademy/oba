## The ECDSA Algorithm

ECDSA (Elliptic Curve Digital Signature Algorithm) was Bitcoin's original signature scheme. It operates on the secp256k1 elliptic curve, the same curve used for key generation.

To sign a message, the algorithm takes the private key and a hash of the message, combines them with a random number (called k, or the nonce), and produces two values: r and s. Together, (r, s) form the signature.

To verify, the algorithm takes the public key, the message hash, and the signature (r, s). It performs elliptic curve math to check whether the signature is consistent with the public key and message. No private key is needed for verification.
