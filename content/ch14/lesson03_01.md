## When Key Path Fails

If the key path cannot be used (a signer is unavailable, a timelock needs to be exercised), the script path provides an alternative. The spender reveals one script from the tree committed in Q, proves it was committed at creation time, and satisfies that script.

The witness for a script path spend contains: the data that satisfies the script (signatures, preimages), the script itself, and a control block. The control block contains the internal key P, a leaf version byte, and the Merkle proof (the sibling hashes along the path from the script leaf to the root).

The verifier reconstructs the Merkle root from the revealed script and proof, computes the expected tweak, and checks that Q equals P + tweak * G. If the math checks out, the script was committed at creation time.
