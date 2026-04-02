## Expanding Multisig with Time

Timelocked policies can expand their signer set over time. A business might start with a 2-of-2 policy between two co-founders. After one year, a third key (held by legal counsel) activates, and the policy becomes 2-of-3.

The Miniscript expression is: `or(multi(2,founder_a,founder_b),and(multi(2,founder_a,founder_b,counsel),older(52560)))`. During the first year, only the two founders can sign. After 52,560 blocks (approximately one year), any two of the three keys can sign.

The entire policy lives in a single UTXO. No on-chain transaction is needed when the timelock expires. The additional spending path was committed at output creation time. The counsel's key gains spending power automatically when the block height passes the threshold. This removes the need for an active key ceremony at the transition point.

Miniscript makes these compositions auditable. The compiler can enumerate every spending path and its conditions. A reviewer can verify that the counsel's key has no spending power before the timelock, that the founders retain full control at all times, and that the witness sizes stay within consensus limits.
