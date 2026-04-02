## The 51% Attack and Double Spending

An attacker who controls more than half the network's hash power can outpace the honest chain. The attacker mines blocks in secret, building a private branch longer than the public one. After the merchant accepts a payment with several confirmations, the attacker releases the private branch. The network reorganizes to the longer chain, and the attacker's version of history replaces the original. The payment disappears.

This is a double-spend attack. The attacker spends the same coins twice: once to the merchant on the public chain, and once back to themselves on the private chain. The cost is enormous. The attacker must sustain majority hash power for the duration of the attack and forfeits the block rewards on the abandoned public chain.

With less than 50% of hash power, the probability of success drops exponentially with each confirmation the defender waits. Six confirmations make a sub-majority attack statistically negligible.
