## What Miniscript Solves

Miniscript, specified in BIP 379, is a structured representation of Bitcoin Script. It maps a defined set of composable fragments to Script opcodes. Each fragment has known properties: its type, resource cost, and required witness data.

Because the mapping is structured, software can analyze a Miniscript expression and determine every spending path, every required key, every timelock constraint, and the exact witness size for each path. Two expressions that encode the same policy can be compared and shown equivalent.

Composition becomes safe. A wallet can take two Miniscript fragments, combine them with AND or OR, and compute the resulting script's resource consumption before broadcasting. If any execution path exceeds consensus limits, the compiler rejects the composition at build time, not at spend time.
