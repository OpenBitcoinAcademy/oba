## Nonce Reuse is Catastrophic

FROST inherits Schnorr's sensitivity to nonce management. If a signer uses the same nonce for two different signing sessions, an attacker who observes both signature shares can extract that signer's secret key share. With enough extracted shares (t of them), the attacker can reconstruct the full group secret key and steal all funds.

This is not a theoretical concern. Deterministic nonce generation, which is safe for single-signer Schnorr, is dangerous in the threshold setting. If a malicious coordinator replays an old signing request, a signer using deterministic nonces would produce a new signature share with the same nonce, leaking their share. FROST therefore requires cryptographically random nonces for every signing session.

Implementations must generate nonces from a strong random source and must never persist nonces across sessions. If a signing session is aborted, the nonces used in that session must be discarded. Re-entering a session with the same nonces is equivalent to nonce reuse.
