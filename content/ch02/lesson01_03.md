## Transactions, Mining, and the Blockchain

A transaction is a data structure that transfers bitcoin from one owner to another. Inputs reference previous outputs being spent. Outputs create new spendable amounts locked to recipient keys.

Mining is the process that adds transactions to the blockchain. Miners compete to solve a computational puzzle (proof of work) by hashing candidate blocks until the result falls below a target value. The first miner to find a valid solution broadcasts the block. Other nodes verify it and accept it.

The blockchain is the chain of all valid blocks, linked by hashes. Each block references the block before it. Changing any past block would require redoing the proof of work for that block and every block after it.
