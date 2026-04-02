## Lagrange Interpolation

FROST builds on Shamir's Secret Sharing. During key generation, the group secret key is split into n shares such that any t shares can reconstruct it. The mathematical tool that makes this possible is Lagrange interpolation.

A polynomial of degree t-1 is uniquely determined by t points. Each signer holds one point (their secret share) on a polynomial whose constant term is the group secret. During signing, the signers never reconstruct the full polynomial. Instead, each signer multiplies their signature share by a Lagrange coefficient that depends on the set of participating signers.

The Lagrange coefficient for signer $i$ in a signing set $S$ is: $\lambda_i = \prod_{j \in S, j \neq i} \frac{j}{j - i}$. These coefficients ensure that the sum of the weighted signature shares produces a valid Schnorr signature for the group public key, without any party ever holding or reconstructing the full secret key.
