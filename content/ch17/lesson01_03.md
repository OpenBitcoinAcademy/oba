## Privacy Through Indistinguishability

The privacy gain from threshold signatures goes beyond hiding the spending policy. With OP_CHECKMULTISIG, the script template itself is a fingerprint. Chain analysis firms categorize addresses by script type, identifying multisig wallets and inferring custody arrangements.

FROST eliminates this fingerprint. A 2-of-3 custody wallet, a 5-of-7 corporate treasury, and a single-key personal wallet all produce identical on-chain outputs. Each spends with a 32-byte public key and a 64-byte signature inside a Taproot key path.

This indistinguishability benefits all Taproot users. The larger the set of transactions that look alike, the harder it is to distinguish any single transaction. Threshold signature users blend into the crowd of single-key spenders, and single-key spenders gain plausible deniability that they might be threshold signers.
