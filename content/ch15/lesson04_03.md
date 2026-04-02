## Products Built on Miniscript

Liana Wallet is an open-source Bitcoin wallet built specifically around Miniscript timelocked recovery. Developed by Wizardsardine and released under the BSD 3-Clause license, Liana lets users configure a primary spending key with one or more recovery paths that activate after chosen delays. The wallet handles descriptor generation, address derivation, and PSBT construction. Users set their policy through a graphical interface without writing Miniscript by hand.

AnchorWatch takes a different approach. Founded by Rob Hamilton, AnchorWatch provides insured Bitcoin custody backed by Lloyd's of London. The custody model uses Miniscript policies where AnchorWatch holds a recovery key that activates only after a timelock. The insurance covers loss of the primary key. Because the spending conditions are encoded in Miniscript, the insurer can verify the policy on-chain: the recovery key has no access before the timelock, and the primary holder retains full control during normal operation.

Both products exist because Miniscript made complex spending policies portable and verifiable. The policy is in the descriptor. Any compatible wallet, signer, or auditor can parse it and confirm the conditions. The trust is in the math, not the vendor.
