## P2PKH Step by Step

Here is what happens when a P2PKH script executes:

1. The signature and public key are pushed onto the stack (from the unlocking script).
2. OP_DUP duplicates the public key on the stack.
3. OP_HASH160 hashes the top element (the duplicated public key) with SHA-256 then RIPEMD-160.
4. The expected public key hash is pushed onto the stack (from the locking script).
5. OP_EQUALVERIFY pops the top two elements and checks they are equal. If they differ, the script fails.
6. OP_CHECKSIG pops the signature and public key, verifies the signature against the transaction data. If valid, it pushes true.

The result: true on the stack. The spend is authorized. The signature proves the spender controls the private key, without revealing it.
