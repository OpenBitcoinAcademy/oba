## Three Layers: Policy, Miniscript, Script

Miniscript operates in a three-layer architecture. At the top is a human-readable policy language. In the middle is Miniscript itself. At the bottom is Bitcoin Script.

A policy describes what you want: "Alice AND Bob, OR Carol after 90 days." You write this as `or(and(pk(Alice),pk(Bob)),and(pk(Carol),older(12960)))`. The policy language is for humans. It names keys and timelocks without concern for opcodes.

A compiler translates the policy into a Miniscript expression: `or_d(and_v(v:pk(Alice),pk(Bob)),and_v(v:pk(Carol),older(12960)))`. The Miniscript expression specifies exactly how the conditions are composed, including which fragment types are used at each position. From there, the expression maps directly to Bitcoin Script opcodes. Each Miniscript fragment has one and only one Script encoding.
