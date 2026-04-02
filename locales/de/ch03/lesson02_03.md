## Anforderungen an Ressourcen

Der Betrieb einer Full Node erfordert Speicherplatz, Bandbreite und Zeit. Der initiale Download der Blockchain umfasst über 500 GB (Stand 2024) und wächst um etwa 150 MB pro Tag.

Bitcoin Core kann im Pruned-Modus laufen: Alte Blockdaten werden nach der Verifizierung verworfen und nur das UTXO-Set und aktuelle Blocks behalten. Das reduziert den Speicherbedarf auf unter 10 GB. Die Node verifiziert trotzdem die gesamte Geschichte beim initialen Sync.

Eine Full Node kann auf Hardware so bescheiden wie einem Raspberry Pi laufen. Die initiale Synchronisierung dauert auf langsamer Hardware länger, aber nach der Synchronisierung ist das tägliche Datenvolumen bei den meisten Heiminternetverbindungen handhabbar.
