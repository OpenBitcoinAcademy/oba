## The Problem Before PSBTs

A Bitcoin transaction needs information from multiple sources before it can be broadcast. The wallet that creates the transaction must know which UTXOs to spend. The signer must hold the private keys. In many setups, these are different devices or different software.

Before BIP 174, there was no standard way to pass an unsigned transaction between these participants. Each wallet software invented its own format. Bitcoin Core serialized partial transactions differently from Electrum, which serialized them differently from hardware wallets. A transaction created in one tool could not be signed by another without custom glue code.

This made multisignature setups painful. Each signer needed compatible software. Coordinating between a laptop, a hardware wallet, and a cold storage machine required manual steps and format conversions that introduced errors.
