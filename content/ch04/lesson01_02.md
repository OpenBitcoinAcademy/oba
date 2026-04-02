## Three Key Properties

Hash functions have three properties that make them useful for Bitcoin.

**Deterministic.** The same input always produces the same output. Hash "bitcoin" a million times and you get the same result every time.

**Avalanche effect.** Change a single bit of the input and the output changes completely. "bitcoin" and "Bitcoin" produce entirely different hashes. There is no way to predict how the output will change.

**One-way.** Given a hash output, there is no way to work backwards to find the input. You can go from input to hash, but not from hash to input. The only option is to guess inputs and check.
