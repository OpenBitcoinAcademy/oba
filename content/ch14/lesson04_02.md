## Removed Limits

Legacy Script and SegWit v0 impose strict resource limits: 10,000-byte script size, 201 non-push opcodes, 100 witness stack elements. These limits were necessary because the cost model did not account for individual scripts.

Tapscript removes the script size limit and the opcode count limit. Instead, it uses a per-input signature operations budget: 50 sigops + the witness weight of the input. Larger witnesses pay higher fees, and the sigops budget scales with the fee paid. The economic incentive replaces the fixed limit.

Scripts that were impossible under the old limits now work. A threshold with 100 participants, a complex timelocked cascade, or a hash chain verification can be expressed in a single Tapscript leaf.
