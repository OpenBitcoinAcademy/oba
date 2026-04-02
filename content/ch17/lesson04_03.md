## The Bigger Picture

Threshold signatures represent a shift in how Bitcoin custody works. The legacy approach revealed spending policies on-chain: anyone could count the keys, see the threshold, and track multisig wallets across transactions. FROST and ChillDKG move all of that complexity off-chain.

The group negotiates its threshold and generates keys using ChillDKG. Any t signers produce a signature using the FROST protocol. The blockchain sees a Taproot key path spend. Verifiers check one signature against one key. The on-chain cost is constant regardless of how many participants are involved: 64 bytes for the signature, 32 bytes for the key.

This is possible because Taproot and BIP 340 Schnorr verification are already deployed on the Bitcoin network. No soft fork is required. No new opcodes. The threshold signing protocol runs entirely in the wallets of the participants. The consensus layer verifies the result using the same rules it uses for every other Taproot spend.
