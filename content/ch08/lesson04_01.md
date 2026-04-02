## SIGHASH Types

A signature does not commit to the entire transaction by default. The SIGHASH flag appended to each signature specifies which parts of the transaction the signature covers.

**SIGHASH_ALL** (0x01) is the default. The signature commits to all inputs and all outputs. Changing any part of the transaction invalidates the signature.

**SIGHASH_NONE** (0x02) commits to all inputs but no outputs. Anyone can change the outputs after signing. Used in rare collaborative protocols.

**SIGHASH_SINGLE** (0x03) commits to all inputs but only the output at the same index as the input being signed. Other outputs can be changed.

The ANYONECANPAY modifier can combine with any of these, allowing signatures that commit to only a single input.
