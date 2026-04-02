## Native SegWit (P2WPKH und P2WSH)

Segregated Witness hat die Signature-Daten aus dem Input Script in eine separate Witness-Struktur verschoben. Das behob Transaction Malleability (Dritte konnten die txid durch Manipulation der Signature-Kodierung ändern) und ermöglichte den Witness-Rabatt bei Gebühren.

P2WPKH (Pay to Witness Public Key Hash) ist das Segwit-Äquivalent zu P2PKH. Das Output Script enthält: OP_0 <20-byte-pubkey-hash>. Der Witness liefert die Signature und den Public Key. Das Input Script ist leer.

P2WSH (Pay to Witness Script Hash) ist das Segwit-Äquivalent zu P2SH. Das Output Script enthält: OP_0 <32-byte-script-hash>. Der Witness liefert das Script und seine benötigten Daten. P2WSH verwendet einen 32-Byte-SHA-256-Hash statt des 20-Byte-HASH160 von P2SH und bietet stärkeren Kollisionsschutz.
