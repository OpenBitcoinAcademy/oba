## Replace By Fee (RBF)

If a transaction is stuck in the mempool because its fee rate is too low, the sender can replace it with a new version that pays a higher fee. This is Replace By Fee (RBF), defined in BIP125.

To use RBF, the original transaction must signal replaceability by setting the sequence number of at least one input to a value below 0xFFFFFFFE. Wallet software handles this automatically.

The replacement transaction must pay a higher total fee than the original. It can change the outputs (paying a different amount or adding/removing recipients) as long as it spends at least one of the same inputs.
