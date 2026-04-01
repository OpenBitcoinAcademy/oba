## From Private Key to Public Key

From a private key, you can calculate a public key. This uses elliptic curve multiplication, a mathematical operation that is easy to perform in one direction and computationally infeasible to reverse.

The private key is a number. Multiply it by a specific point on a known curve (called the generator point G), and you get another point on the curve. That result is your public key.

Anyone can verify that a public key matches a private key by checking the math. But given only the public key, no one can determine the private key. This one-way relationship is the foundation of Bitcoin's security.
