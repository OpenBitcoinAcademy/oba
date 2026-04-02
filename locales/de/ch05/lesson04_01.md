## BIP39 Seed Phrases

BIP39 kodiert einen Wallet-Seed als Folge von 12 oder 24 englischen Wörtern, die Recovery Phrase (oder Seed Phrase) genannt wird. Jedes Wort stammt aus einer standardisierten Liste von 2.048 Wörtern.

Die Wort-Kodierung hat zwei Zwecke. Wörter lassen sich leichter aufschreiben, vorlesen und abschreiben als rohe Hexadezimalwerte. Das letzte Wort enthält eine Prüfsumme, die Abschreibfehler erkennt.

Eine 12-Wort-Phrase kodiert 128 Bit Entropie. Eine 24-Wort-Phrase kodiert 256 Bit. Beide sind stark genug für aktuelle Sicherheitsanforderungen. Die Phrase wird mit PBKDF2 und 2.048 Runden HMAC-SHA512 in einen 512-Bit-Seed umgewandelt.
