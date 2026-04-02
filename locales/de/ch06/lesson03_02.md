## Byte für Byte

**Version** (4 Bytes): 01000000 bedeutet Version 1. 02000000 bedeutet Version 2 (aktiviert BIP68-Sequence-Einschränkungen). Little-Endian-Kodierung: Das niederwertigste Byte steht zuerst.

**Inputs**: Eine Varint-Anzahl gefolgt von jedem Input. Jeder Input enthält den vorherigen Transaction-Hash (32 Bytes, umgekehrt), den Output-Index (4 Bytes), die Input-Script-Länge (Varint), das Input Script (variabel) und die Sequence Number (4 Bytes).

**Outputs**: Eine Varint-Anzahl gefolgt von jedem Output. Jeder Output enthält den Wert in Satoshis (8 Bytes, Little-Endian) und die Output-Script-Länge (Varint) gefolgt vom Output Script.

**Witness** (nur Segwit): Für jeden Input eine Anzahl von Witness-Elementen gefolgt von Länge und Daten jedes Elements. Legacy-Inputs haben null Witness-Elemente.

**Locktime** (4 Bytes): Normalerweise 00000000. Ein Wert ungleich null schränkt ein, wann die Transaction gemined werden kann.

Die Transaction-ID (txid) ist der doppelte SHA-256-Hash der serialisierten Transaction, wobei die Witness-Daten ausgeschlossen (bei Segwit) oder eingeschlossen (bei Legacy) werden.
