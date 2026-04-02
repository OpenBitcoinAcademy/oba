## Varints und Endianness

Bitcoin verwendet zwei Kodierungskonventionen in seiner gesamten Serialisierung.

**Varints** (variable Ganzzahlen) kodieren Anzahlen und Längen. Ein Wert unter 253 passt in ein Byte. Werte von 253-65535 verwenden ein 0xFD-Präfix gefolgt von 2 Bytes. Werte bis ~4 Milliarden verwenden 0xFE gefolgt von 4 Bytes. Größere Werte verwenden 0xFF gefolgt von 8 Bytes.

**Little-Endian**-Byte-Reihenfolge stellt das niederwertigste Byte an den Anfang. Version 1 wird als 01 00 00 00 serialisiert, nicht als 00 00 00 01. Satoshi-Werte, Blockhöhen und die meisten numerischen Felder verwenden Little-Endian.

Transaction-Hashes (txids) werden konventionsgemäß in umgekehrter Byte-Reihenfolge angezeigt. Der doppelte SHA-256-Hash erzeugt Bytes in einer Reihenfolge, aber Block Explorer und Bitcoin Core zeigen sie umgekehrt an. Das ist eine Anzeigekonvention, keine Serialisierungsregel.
