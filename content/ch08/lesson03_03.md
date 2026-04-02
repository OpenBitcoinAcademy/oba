## Schnorr vs ECDSA

Both algorithms use the secp256k1 curve. Both provide equivalent security against key recovery. The differences are practical.

Schnorr signatures are smaller (64 bytes vs 71-73). Schnorr supports native signature aggregation. Schnorr has a simpler verification equation. Schnorr has a formal security proof.

ECDSA was chosen for Bitcoin in 2008 because Schnorr was patented until 2008, and the patent status was uncertain when Satoshi designed the system. With the patent expired and Taproot activated, new transactions can use Schnorr. ECDSA remains supported for backward compatibility.
