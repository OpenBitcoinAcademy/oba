## The JSON-RPC Interface

Bitcoin Core exposes a JSON-RPC interface that allows programs to query blockchain data and submit transactions. The command-line tool `bitcoin-cli` sends requests to this interface.

Every piece of data in the blockchain is queryable: block headers, full blocks, individual transactions, address balances, mempool contents, peer connections, and network statistics.

Wallets, block explorers, payment processors, and research tools all communicate with Bitcoin Core through this API. The interface is the bridge between the raw blockchain data and the applications that present it to users.
