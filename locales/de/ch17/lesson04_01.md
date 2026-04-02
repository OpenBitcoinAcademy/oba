## Nonce-Wiederverwendung ist katastrophal

FROST erbt Schnorrs Empfindlichkeit gegenüber dem Nonce-Management. Wenn ein Unterzeichner dieselbe Nonce für zwei verschiedene Signiersitzungen verwendet, kann ein Angreifer, der beide Signature Shares beobachtet, den Secret Key Share dieses Unterzeichners extrahieren. Mit genügend extrahierten Shares (t davon) kann der Angreifer den vollständigen Group Secret Key rekonstruieren und alle Mittel stehlen.

Das ist keine theoretische Sorge. Deterministische Nonce-Erzeugung, die für Single-Signer-Schnorr sicher ist, ist im Threshold-Setting gefährlich. Wenn ein bösartiger Coordinator eine alte Signieranfrage wiederholt, würde ein Unterzeichner mit deterministischen Nonces einen neuen Signature Share mit derselben Nonce erzeugen und so seinen Share preisgeben. FROST verlangt daher kryptographisch zufällige Nonces für jede Signiersitzung.

Implementierungen müssen Nonces aus einer starken Zufallsquelle erzeugen und dürfen Nonces nie über Sitzungen hinweg speichern. Wenn eine Signiersitzung abgebrochen wird, müssen die in dieser Sitzung verwendeten Nonces verworfen werden. Das Wiedereintreten in eine Sitzung mit denselben Nonces ist gleichbedeutend mit Nonce-Wiederverwendung.
