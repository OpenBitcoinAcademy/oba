## Public Child Key Derivation

One feature of HD wallets is the ability to derive public child keys from a parent public key, without knowing the parent private key.

A web server generating receiving addresses for an online store can hold only the parent public key. It derives fresh receiving addresses for each customer without having access to any private key. The private keys needed to spend funds exist only on a separate, secured device.

This separation of public and private key spaces is a major security advantage. The online system that handles customer-facing operations never touches spending keys.
