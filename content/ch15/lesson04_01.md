## Time-Decayed Spending Policies

A timelocked recovery policy starts with a primary spending condition and adds fallback conditions that activate after a delay. The primary key controls the funds immediately. If the primary key is lost or its holder becomes unavailable, an alternative spending path unlocks after a specified number of blocks.

In Miniscript, this is expressed as: `or(pk(primary),and(pk(recovery),older(26280)))`. The primary key can spend at any time. The recovery key can spend only after approximately six months (26,280 blocks at 10 minutes per block). The policy compiles to a single Script with two execution paths.

This pattern solves a real problem. A single-key wallet has no recovery path. If the key is lost, the funds are gone. A standard multisig requires multiple signers for every transaction, adding friction to daily use. A timelocked fallback gives the primary holder full control during normal operation while guaranteeing that a trusted party can recover funds after a delay.
