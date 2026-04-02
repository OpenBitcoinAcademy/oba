## Transaction Pinning

Transaction pinning is an attack where a malicious party creates a low-fee child transaction that makes it expensive or impossible for the honest party to fee-bump the parent.

In a two-party protocol (like a Lightning Network channel), one party might broadcast a large, low-fee descendant of a shared transaction. The other party's CPFP attempt would need to pay for the attacker's large descendant, making the fee bump prohibitively expensive.

Anchor outputs are a countermeasure. Each party gets a small output in the shared transaction that they can spend with CPFP. Rules limit how many descendants each anchor can have, preventing the pinning attack.
