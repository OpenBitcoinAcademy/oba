## Cashu: Single-Mint Ecash

Cashu ist ein Chaumian-Ecash-Protokoll mit einem einzelnen Mint-Betreiber anstelle einer Federation. Nutzer zahlen Bitcoin ein (typischerweise per Lightning) und erhalten blind signierte Ecash-Tokens. Die Tokens sind Inhaberinstrumente: Wer ein gültiges Token hält, kann es einlösen.

Die Mint kann Minting nicht mit Einlösung verknüpfen (Blind Signatures stellen das sicher). Die Mint kann nicht sehen, wer Tokens an wen überträgt. Transfers zwischen Nutzern sind sofort und kostenlos (keine Blockchain-Interaktion).

Der Kompromiss: Nutzer vertrauen einem einzelnen Mint-Betreiber. Wenn die Mint unehrlich ist oder offline geht, können Mittel verloren gehen. Cashu ist für kleine, anwendungsspezifische Deployments konzipiert (ein Cafe, ein Streaming-Dienst, eine Community), bei denen die Vertrauensbeziehung handhabbar ist.
