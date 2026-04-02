## X-Only Public Keys

Taproot verwendet 32-Byte X-only Public Keys anstelle des traditionellen 33-Byte komprimierten Formats. Die Y-Koordinate wird implizit als gerade gewählt. Das spart ein Byte pro Key on-chain und in Signatures.

Ein komprimierter Standard-Public-Key hat ein Präfix-Byte (0x02 für gerade Y, 0x03 für ungerade Y), gefolgt von 32 Bytes X-Koordinate. Taproot lässt das Präfix weg und verlangt, dass die Y-Koordinate gerade ist. Hat das berechnete Q ein ungerades Y, negiert die Implementierung den Private Key (was Y auf gerade kippt), bevor signiert wird.

Diese Konvention vereinfacht Batch-Verifikation und Key-Aggregation. Jeder Taproot Output Key ist 32 Bytes. Jede Schnorr Signature ist 64 Bytes. Es gibt keine Kodierungen variabler Länge.
