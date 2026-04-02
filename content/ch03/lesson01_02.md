## From Client to Platform

The original Bitcoin software included a wallet, a miner, and a full node in a single program. Over time, these functions separated. Mining moved to specialized hardware and software. Wallet functions moved to dedicated applications.

Bitcoin Core evolved into a platform for full node operation. It downloads and verifies the entire blockchain, maintains the UTXO database, relays transactions and blocks, and exposes an API for querying blockchain data.

Most full node operators on the Bitcoin network run Bitcoin Core. It sets the reference behavior that other implementations measure themselves against.
