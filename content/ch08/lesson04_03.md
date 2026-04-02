## SegWit's Signing Algorithm

Segwit introduced a new signing algorithm (BIP143) that fixes a longstanding problem: in legacy transactions, the amount being spent is not included in the signed data. This forced signers to fetch the full previous transaction to verify the amount, slowing down hardware wallets.

BIP143 includes the amount of each input in the data that is signed. A hardware wallet can verify the total input value and the fee without downloading previous transactions. The signing process is faster and more secure.

For segwit v1 (Taproot), BIP341 defines an updated algorithm that includes additional commitments, supporting both key path and script path spending.
