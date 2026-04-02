## Constructing the Transaction

A legacy transaction has four fields: version, inputs, outputs, and locktime. Modern segwit transactions add three more: a marker, a flag, and a witness structure that holds the authorization data (signatures) separately from the inputs.

**Version** is a number (currently 1 or 2) that tells nodes which validation rules apply.

**Inputs** list the UTXOs being spent. Each input specifies the previous transaction ID, the output index within that transaction, an input script, and a sequence number.

**Outputs** list the new UTXOs being created. Each output specifies an amount in satoshis and a locking script.

**Locktime** is usually zero. When set to a future block height or timestamp, the transaction cannot be included in a block until that point.

The transaction is serialized into a byte sequence, hashed twice with SHA-256, and the result is the transaction ID.
