## MAST and Tapscript

The Merkle tree in Taproot is a Merklized Alternative Script Tree (MAST). Each leaf of the tree is a different script (a different spending condition). To spend via script path, the spender reveals only the script they are using and provides a Merkle proof that it is part of the committed tree.

Scripts not used remain hidden. A Taproot output with 100 possible spending conditions looks the same on the blockchain as one with 1 condition, unless the script path is exercised. Unused branches stay private.

Tapscript (BIP342) is the script language used inside Taproot's script-path leaves. It is similar to legacy Script but with improvements: OP_CHECKSIGADD replaces OP_CHECKMULTISIG (enabling batch validation), signature opcodes use Schnorr instead of ECDSA, and the script size limit is removed (replaced by weight limits).
