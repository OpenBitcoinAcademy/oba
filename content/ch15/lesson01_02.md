## Resource Limits and Hidden Failures

Bitcoin consensus enforces resource limits on scripts: a maximum size of 10,000 bytes in legacy, a 201 non-push opcode limit, and a sigops budget. A script that exceeds any limit is invalid. The transaction carrying it will never confirm.

When composing scripts by hand, resource consumption is hard to predict. Each branch of an OP_IF tree contributes differently to the opcode count. Nested conditions multiply complexity. A script might work in one branch and exceed the opcode limit in another, depending on which execution path a spender takes.

These failures are silent. The script looks correct in a text editor. It might pass unit tests for one spending path. The failing path is discovered only when someone tries to use it on the network, and the transaction is rejected.
