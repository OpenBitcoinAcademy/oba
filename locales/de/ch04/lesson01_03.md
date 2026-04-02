## SHA-256 in Bitcoin

Bitcoin verwendet eine Hash-Funktion namens SHA-256, entworfen von der National Security Agency der Vereinigten Staaten. SHA steht für Secure Hash Algorithm. 256 bezieht sich auf die Ausgabegröße: 256 Bit oder 32 Byte.

SHA-256 kommt überall in Bitcoin vor. Transaction-IDs sind Double-SHA-256-Hashes. Mining besteht darin, Inputs zu finden, die Hashes unterhalb eines Zielwerts erzeugen. Die Address-Ableitung nutzt SHA-256 in Kombination mit RIPEMD-160 (einer anderen Hash-Funktion, die einen kürzeren 20-Byte-Output erzeugt). Die Kombination beider, HASH160 genannt, erzeugt den Public Key Hash in einer Address.
