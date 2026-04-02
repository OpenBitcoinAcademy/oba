## BitVM: Computation Verification

BitVM enables Turing-complete computation verification on Bitcoin without consensus changes. A prover claims that a function f(x) = y. If the claim is correct, nothing happens on-chain. If incorrect, any verifier can submit a fraud proof and the prover loses a bond.

BitVM2 (2024-2025) made verification permissionless: anyone can challenge a false claim, and disputes resolve in three on-chain transactions. The system implements a SNARK verifier in Bitcoin Script, enabling efficient verification of complex computations.

Applications include trustless bridges (moving bitcoin between chains), rollups (batching transactions off-chain with on-chain verification), and any protocol that benefits from "verify, don't compute" on Bitcoin.
