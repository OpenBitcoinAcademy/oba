## The Taproot Safety Tweak

ChillDKG produces a group public key, but this key cannot be used directly as a Taproot output key. BIP 341 requires the output key Q to be a tweaked version of an internal key P: Q = P + tagged_hash("TapTweak", P) * G for a key-path-only output.

BIP 445 specifies that ChillDKG applies this tweak as part of key generation. The protocol computes the tweaked group public key and adjusts the key shares so that the group can sign for the tweaked key without any additional steps during signing. Each participant's share is modified to account for the tweak, and the group public key returned by ChillDKG is the final Taproot output key.

This integration matters because the tweak must be applied consistently. If participants disagree on the tweaked key, signing will fail. By embedding the tweak into the DKG protocol itself, ChillDKG guarantees that all participants derive the same output key and hold shares that are consistent with it.
