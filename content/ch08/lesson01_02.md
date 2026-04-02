## Creating and Verifying

Creating a signature requires two inputs: the private key and the message (the transaction data, or a hash of it). The signature algorithm produces a signature value that is unique to both the key and the message.

Verifying a signature requires three inputs: the public key, the message, and the signature. The verification algorithm outputs true or false. If true, the signature was produced by the holder of the corresponding private key for that exact message.

The private key is never revealed during creation or verification. Anyone can verify, but only the key holder can sign.
