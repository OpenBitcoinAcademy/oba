## Three Layers: SimplPedPop, EncPedPop, ChillDKG

The foundation is SimplPedPop. Each participant generates their own random polynomial of degree t-1 and sends a secret evaluation to every other participant, along with a commitment to their polynomial's coefficients. Every participant sums the evaluations they received to compute their final secret share. The group public key is derived from the sum of all participants' commitments to their constant terms. No party ever holds the full secret.

SimplPedPop assumes a secure channel between each pair of participants. EncPedPop adds this by having each participant generate an ephemeral encryption key pair and encrypting their secret evaluations with the recipient's public key. Now the protocol works over a public broadcast channel, because eavesdroppers cannot decrypt the secret shares.

ChillDKG wraps EncPedPop with a session management layer. It introduces a host secret key that each participant holds persistently, a common recovery dataset that is identical for all participants, and a protocol for detecting and handling misbehavior. The host secret key, combined with the common recovery data, allows a participant to reconstruct their share if they lose their signing device.
