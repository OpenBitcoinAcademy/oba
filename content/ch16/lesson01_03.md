## Why Hardware Wallets Needed This

Hardware wallets have a specific constraint: they cannot query the blockchain. They receive data, verify what they can, sign, and return the result. Without a standard format, every hardware wallet manufacturer had to write custom integrations for every software wallet. Adding a new hardware wallet meant patching every coordinator. Adding a new coordinator meant patching every hardware wallet. The combinatorial explosion was unsustainable.

PSBT solved this by defining one format that carries all the information a signer needs. The UTXO being spent, the derivation path for the key, the sighash type to use, the redeem script for P2SH, the witness script for P2WSH. A hardware wallet receives a PSBT, reads the fields it needs, signs, writes its partial signature back into the PSBT, and returns it. No proprietary protocol. No format negotiation.

The ecosystem converged quickly. Coldcard, Trezor, Ledger, BitBox, and Jade all speak PSBT. Software coordinators like Sparrow, Specter, and Bitcoin Core all produce and consume PSBTs. A multisig quorum can use hardware from different vendors without compatibility worries.
