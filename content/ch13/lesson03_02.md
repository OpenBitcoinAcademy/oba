## Consensus Risks

Applications that validate transactions must implement the consensus rules exactly. A subtle difference in how two implementations handle an edge case can cause them to disagree on the validity of a block, splitting the network from their perspective.

The safest approach for applications: delegate consensus validation to a full node (Bitcoin Core or an equivalent) and use its API for blockchain queries. Do not reimplement consensus rules in application code unless you are prepared to match every edge case in the reference implementation.

For wallet developers: use well-tested libraries for key derivation, address generation, and transaction signing. Prefer battle-tested open-source implementations over custom code.
