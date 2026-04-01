## From Public Key to Address

You do not give people your public key directly. Instead, you give them an address: a shorter, checksummed string derived from your public key.

An address serves two purposes. It shortens the public key from 33 bytes (compressed) to 20 bytes. And it adds a checksum that catches typos. If you mistype a single character, the address becomes invalid, and no wallet will send to it.

Think of the address as a mailbox. Anyone can drop bitcoin into it using the address. Only the person with the matching private key can take bitcoin out.
