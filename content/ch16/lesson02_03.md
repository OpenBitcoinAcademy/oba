## Combiner, Finalizer, Extractor

The Combiner merges multiple PSBTs that contain partial signatures for the same transaction. In a 2-of-3 multisig, signer A produces a PSBT with its signature and signer B produces another. The Combiner takes both, merges the PARTIAL_SIG entries for each input, and produces a single PSBT with all available signatures.

The Finalizer transforms partial signatures into a complete scriptSig and witness for each input. For a P2PKH input, it takes the single PARTIAL_SIG and builds the scriptSig. For a P2WSH 2-of-3 multisig, it takes the partial signatures, orders them correctly, and constructs the witness stack with the redeem script. After finalization, the PSBT input maps contain FINAL_SCRIPTSIG and FINAL_SCRIPTWITNESS fields. The partial data is no longer needed.

The Extractor reads the finalized PSBT and produces the raw network transaction. It takes the unsigned transaction from the global map, inserts the FINAL_SCRIPTSIG and FINAL_SCRIPTWITNESS from each input, and serializes the result. The output is a standard Bitcoin transaction ready for broadcast. The PSBT has served its purpose.
