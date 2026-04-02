## What a Transaction Looks Like

A Bitcoin transaction is a data structure that transfers bitcoin from one owner to another. It contains inputs (which reference previous outputs to spend), outputs (which create new spendable amounts), and authorization data (signatures proving the spender controls the keys).

Inputs say where the bitcoin comes from. Each input references a previous transaction's output that has not been spent yet. This unspent output is called a UTXO (unspent transaction output).

Outputs say where the bitcoin goes. Each output specifies an amount in satoshis and a locking condition (a script) that determines who can spend it.

A transaction consumes old outputs and creates new ones. Nothing is "stored in an account." Bitcoin tracks ownership through a chain of outputs, each locked to a specific key.
