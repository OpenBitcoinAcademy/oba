## Ausgeben mit dem getweakten Key

Key Path Spending ist die einfachste und privateste Art, einen Taproot Output auszugeben. Die Witness enthält ein einziges Element: eine 64-Byte BIP 340 Schnorr Signature für den getweakten Key Q.

Der Absender berechnet den getweakten Private Key (originaler Private Key + Tweak), signiert die Transaction und liefert die Signature. Keine Scripts werden offengelegt. Keine Merkle Proofs werden benötigt. Die gesamte Witness ist 64 Bytes.

Ein Verifizierer prüft die Signature gegen Q (den Key im scriptPubKey des Outputs). Ist sie gültig, ist die Ausgabe autorisiert. Der Verifizierer weiß und interessiert sich nicht dafür, ob Q aus einem einzelnen Key, einem MuSig2-Aggregat oder einem Key mit eingebettetem Script-Baum abgeleitet wurde.
