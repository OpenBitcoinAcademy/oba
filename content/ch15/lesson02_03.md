## Bijective Mapping

The relationship between Miniscript and Script is bijective for the supported fragment set. Each Miniscript expression maps to exactly one Script encoding, and each supported Script pattern maps back to exactly one Miniscript expression.

Consider the fragment `pk(K)`: it encodes to `<K> OP_CHECKSIG`. The fragment `and_v(v:pk(A),pk(B))` encodes to `<A> OP_CHECKSIGVERIFY <B> OP_CHECKSIG`. There is no ambiguity. Given the Script, you can recover the original Miniscript.

This property makes round-trip analysis possible. A wallet can receive a Script from another party, decode it to Miniscript, and determine the spending conditions without trusting the sender's description. Two wallets built by different teams can agree on a policy, compile to Miniscript independently, and verify they produced the same Script.

Not all valid Bitcoin Scripts can be represented in Miniscript. The fragment set covers the patterns needed for practical spending policies: keys, hashes, timelocks, thresholds, and Boolean combinations. Scripts that use unusual opcode sequences or stack manipulation fall outside the Miniscript subset.
