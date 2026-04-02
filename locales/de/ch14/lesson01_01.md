## Der Internal Key und der Tweak

Ein Taproot Output sperrt Mittel auf einen getweakten Public Key Q. Dieser Key wird aus zwei Eingaben abgeleitet: einem internen Public Key P und einem optionalen Merkle Root eines Script-Baums.

Der Tweak-Wert t wird berechnet als: t = tagged_hash("TapTweak", P || merkle_root). Der getweakte Key ist: Q = P + t * G, wobei G der Generatorpunkt auf secp256k1 ist.

Gibt es keinen Script-Baum, verwendet der Tweak nur P: t = tagged_hash("TapTweak", P). Der Output committet trotzdem auf den Internal Key, aber es sind keine Scripts eingebettet.

On-chain lautet der scriptPubKey: OP_1 gefolgt von der 32-Byte-X-Koordinate von Q. Das sind 34 Bytes, dieselbe Größe wie ein P2WSH Output. Kein Beobachter kann erkennen, ob Q einen eingebetteten Script-Baum enthält oder nicht.
