## Child Pays for Parent (CPFP)

CPFP is a fee-bumping technique where the recipient creates a new transaction spending the unconfirmed output. The new transaction (the child) pays a high fee rate. Miners who want the child's fee must include the parent too.

CPFP does not require the original sender's cooperation. The recipient creates the child transaction using their own keys. This makes CPFP useful when the sender is unavailable or unresponsive.

The combined fee rate of parent and child must be attractive enough for a miner to include both. Miners evaluate transaction "packages" (groups of dependent transactions) when selecting which to include.
