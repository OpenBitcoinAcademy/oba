## Single-Use Seals

RGB defines a "seal" as a Bitcoin UTXO. When a state transition occurs (transferring a token, updating a contract), the seal is "closed" by spending that UTXO. A UTXO can be spent only once, so a seal can be closed only once. This prevents double-spending of RGB assets using Bitcoin's existing consensus mechanism.

The state transition data (who owns what, contract updates) never touches the blockchain. Only a cryptographic commitment to the transition is embedded in the Bitcoin transaction, typically in a Taproot output. The blockchain anchors the timing and ordering. The data lives off-chain with the parties involved.
