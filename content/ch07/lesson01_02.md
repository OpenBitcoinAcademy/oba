## Unlocking Scripts

To spend a locked output, the spender provides an unlocking script, called a scriptSig. For P2PKH, the unlocking script contains a digital signature and the spender's public key.

Bitcoin executes these scripts on a stack machine. First, the unlocking script runs and pushes data onto the stack. Then the stack is copied, and the locking script runs against that copied stack. The two scripts never combine into one. This separation was introduced in 2010 to fix a security vulnerability.

If the locking script finishes with a truthy value on top of the stack, the spend is valid. If the stack is empty or the top value is zero, the spend fails.
