## Four Consensus Processes

Bitcoin's consensus is not a single mechanism. It is the result of four independent processes running on every node in the network.

First, each node independently verifies every transaction against a set of rules (valid signatures, unspent inputs, correct amounts). Second, miners aggregate verified transactions into candidate blocks. Third, each node independently validates new blocks against the consensus rules (correct proof of work, valid coinbase, proper structure). Fourth, each node independently selects the chain with the most cumulative proof of work.

No single step is "the consensus." All four processes, running on thousands of independent nodes, produce emergent agreement. No node trusts another. Every node verifies everything for itself.
