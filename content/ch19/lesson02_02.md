## Client-Side Validation

In RGB, only the parties to a contract validate its state transitions. A token issuer and the token holders verify the history of transfers. No other node on the network needs to process or store this data.

Compare this to systems where every node validates every contract (Ethereum's model). Client-side validation scales without limit: adding more contracts does not increase the load on any node that is not involved. Privacy is inherent: contract data is visible only to participants.

The cost: participants must store and verify the full history of their assets. If the history is lost, the asset cannot be verified. RGB uses commitment chains and proofs to make this verification efficient.
