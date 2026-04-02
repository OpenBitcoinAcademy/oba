## Fedimint: Community Custody

Fedimint is a protocol for community-operated custody using federated Chaumian ecash. A group of guardians (trusted community members) jointly operate a mint using threshold multisig. Users deposit bitcoin and receive ecash notes in return.

The ecash notes are created using blind signatures. When minting, the user blinds a random token identifier before sending it to the guardians. The guardians sign it without seeing the identifier. When the user unblinds the signed token, the mint cannot link the deposit to the redemption. Transactions between users are unlinkable.

Users trust the federation's threshold (a majority of guardians must collude to steal). This is weaker than Bitcoin's trustless model but far stronger than trusting a single custodian. Fedimint targets communities with existing social trust: savings groups, cooperatives, local organizations.
