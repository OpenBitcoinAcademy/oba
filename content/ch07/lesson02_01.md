## Pay to Script Hash (P2SH)

P2SH separates the script from the address. Instead of encoding the full locking script in the output, P2SH encodes only a hash of the script. The spender reveals the full script (called the redeem script) when spending.

The output script contains: OP_HASH160 <script_hash> OP_EQUAL. To spend, the input provides the redeem script and whatever data the redeem script requires (signatures, public keys). The network hashes the provided redeem script and checks that it matches the hash in the output.

P2SH addresses start with "3" on mainnet. They enable complex spending conditions (multisig, timelocks) without requiring the sender to know the details. The sender pays to a short hash. The receiver handles the complexity.
