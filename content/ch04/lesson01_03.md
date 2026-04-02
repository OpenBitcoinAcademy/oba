## SHA-256 in Bitcoin

Bitcoin uses a hash function called SHA-256, designed by the United States National Security Agency. SHA stands for Secure Hash Algorithm. 256 refers to the output size: 256 bits, or 32 bytes.

SHA-256 appears everywhere in Bitcoin. Transaction IDs are double-SHA-256 hashes. Mining involves finding inputs that produce hashes below a target value. Address derivation uses SHA-256 combined with RIPEMD-160 (a different hash function that produces a shorter 20-byte output). The combination of both, called HASH160, is what creates the public key hash in an address.
