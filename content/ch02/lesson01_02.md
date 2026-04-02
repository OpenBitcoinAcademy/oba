## Wallets, Keys, and Addresses

A Bitcoin wallet is a collection of keys, not a container of coins. The wallet software manages private keys and uses them to sign transactions. The blockchain records which outputs are spendable by which keys.

To receive bitcoin, you share an address. An address is derived from your public key. Anyone can send bitcoin to your address, but only you can spend it, because spending requires a signature from the corresponding private key.

To send bitcoin, your wallet selects unspent outputs (UTXOs) that you control, constructs a transaction, signs it with your private key, and broadcasts it to the network.
