## Writing Bitcoin Script by Hand

Bitcoin Script is a stack-based language. Each spending condition is a sequence of opcodes: push a key, check a signature, verify a hash, enforce a timelock. For a single-key spend, the script is short and well-understood.

Combine conditions, and the difficulty grows fast. A 2-of-3 multisig with a timelock fallback requires careful ordering of OP_IF branches, OP_CHECKMULTISIG, and OP_CHECKSEQUENCEVERIFY. One misplaced opcode changes the spending conditions entirely. A duplicated key might go unnoticed until funds are locked.

No tool can analyze an arbitrary Script and answer: "Who can spend this, under what conditions?" The language lacks structure. Different opcode sequences can encode the same policy, and there is no general method to compare them.
