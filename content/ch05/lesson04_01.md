## BIP39 Seed Phrases

BIP39 encodes a wallet seed as a sequence of 12 or 24 English words, called a recovery phrase (or seed phrase). Each word comes from a standardized list of 2,048 words.

The word encoding serves two purposes. Words are easier to write down, read back, and transcribe than raw hexadecimal. And the last word includes a checksum that catches transcription errors.

A 12-word phrase encodes 128 bits of entropy. A 24-word phrase encodes 256 bits. Both are strong enough for current security needs. The phrase is converted to a 512-bit seed using PBKDF2 with 2,048 rounds of HMAC-SHA512.
