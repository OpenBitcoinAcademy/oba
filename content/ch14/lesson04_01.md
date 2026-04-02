## A New Script Language

Tapscript (BIP 342) defines the validation rules for scripts executed inside Taproot script path spends. It is similar to legacy Bitcoin Script but with targeted improvements.

All signature-checking opcodes (OP_CHECKSIG, OP_CHECKSIGVERIFY) use Schnorr signatures instead of ECDSA. A new opcode, OP_CHECKSIGADD, replaces the legacy OP_CHECKMULTISIG. Instead of checking multiple signatures against multiple keys in one operation, OP_CHECKSIGADD checks one signature at a time and accumulates a counter. This enables batch verification: the verifier collects all signatures and verifies them together in a single operation, faster than checking each individually.
