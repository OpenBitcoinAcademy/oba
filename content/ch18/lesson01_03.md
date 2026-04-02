## The Revocation Mechanism

When Alice and Bob update their channel balance, they invalidate the previous commitment by exchanging revocation secrets. Each party now holds a secret that can punish the other for broadcasting the old state.

If Bob broadcasts an old commitment where he had 0.5 BTC (but the current balance gives him only 0.3 BTC), Alice can use Bob's revocation secret to claim the entire channel balance. Bob loses everything.

This penalty makes cheating economically irrational. The cost of attempting fraud (losing all funds) far exceeds the potential gain (claiming a slightly higher old balance). Honest behavior is the dominant strategy.
