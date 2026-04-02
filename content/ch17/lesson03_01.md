## The Trusted Dealer Problem

The simplest way to set up FROST key shares is with a trusted dealer. One party generates the group secret key, evaluates the Shamir polynomial at n points, and distributes one share to each participant. This works but creates a single point of failure: the dealer knows the full secret key. If the dealer is compromised, the entire group's funds are at risk.

A compromised dealer can also distribute inconsistent shares, giving some participants invalid data that will cause signing to fail unpredictably. Participants have no way to verify their shares are correct without a protocol that enforces consistency.

ChillDKG, specified in BIP 445, eliminates the trusted dealer. It is a distributed key generation protocol: the group jointly produces the key shares with no single party ever holding or seeing the complete secret key. The protocol builds in three layers, each adding a specific guarantee.
