## Inputs Reference Previous Outputs

Every input points to a specific output from a previous transaction. It does this with two pieces of information: the transaction ID (a hash) and the output index (which output in that transaction).

When you spend bitcoin, you prove that you control the key that can unlock a previous output. In legacy transactions, the proof (signature and public key) lives in the input script. In modern segwit transactions, the proof lives in a separate witness structure, and the input script is empty.

Once an output is referenced by an input, it is spent. It cannot be spent again. This is how Bitcoin prevents double spending without a central authority.
