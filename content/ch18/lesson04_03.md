## The Penalty Close

If a party broadcasts a revoked (old) commitment transaction, the other party can use the revocation secret to claim the entire channel balance. The cheater loses all funds in the channel.

This is the economic enforcement mechanism. Broadcasting old state is detectable (the other party monitors the blockchain for revoked commitments) and punishable (total loss of funds). The penalty makes Lightning channels trustworthy without requiring trust between the parties.

The honest party must be online (or have a watchtower service monitoring on their behalf) to detect and respond to a revoked commitment broadcast within the timelock window.
