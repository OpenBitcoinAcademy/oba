## The Wallet Database

A wallet application stores its data in a wallet database. This database holds private keys, public keys, and metadata like transaction labels, address notes, and fee settings.

Some wallets store only public keys and rely on external devices (hardware wallets) to hold private keys and produce signatures. Others store all keys locally. The security model differs, but the principle is the same: the wallet database is the link between the user and the blockchain.

Wallet applications provide a user interface on top of the database: showing balances, constructing transactions, managing addresses, and connecting to the Bitcoin network.
