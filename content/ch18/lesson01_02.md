## Commitment Transactions

Each party holds their own version of the latest commitment transaction. These transactions split the channel's funds according to the current balance. If Alice has 0.7 BTC and Bob has 0.3 BTC, both commitment transactions reflect this split.

Commitment transactions are asymmetric. Alice's version pays Bob immediately but locks Alice's funds behind a timelock. Bob's version does the opposite. This asymmetry enables the penalty mechanism: if Alice broadcasts an old commitment (trying to claim more than she owns), Bob can use a revocation secret to take all funds before Alice's timelock expires.

Each time the balance updates, both parties exchange revocation secrets for the old commitments, making them unsafe to broadcast.
