## Combiner, Finalizer, Extractor

Der Combiner führt mehrere PSBTs zusammen, die partielle Signatures für dieselbe Transaction enthalten. Bei einem 2-of-3 Multisig erzeugt Signer A eine PSBT mit seiner Signature und Signer B eine weitere. Der Combiner nimmt beide, führt die PARTIAL_SIG-Einträge für jeden Input zusammen und erzeugt eine einzelne PSBT mit allen verfügbaren Signatures.

Der Finalizer transformiert partielle Signatures in ein vollständiges scriptSig und eine Witness für jeden Input. Bei einem P2PKH-Input nimmt er den einzelnen PARTIAL_SIG und baut das scriptSig. Bei einem P2WSH 2-of-3 Multisig nimmt er die partiellen Signatures, ordnet sie korrekt an und konstruiert den Witness Stack mit dem Redeem Script. Nach der Finalisierung enthalten die PSBT-Input-Maps FINAL_SCRIPTSIG- und FINAL_SCRIPTWITNESS-Felder. Die partiellen Daten werden nicht mehr benötigt.

Der Extractor liest die finalisierte PSBT und erzeugt die rohe Netzwerk-Transaction. Er nimmt die unsignierte Transaction aus der globalen Map, setzt FINAL_SCRIPTSIG und FINAL_SCRIPTWITNESS aus jedem Input ein und serialisiert das Ergebnis. Die Ausgabe ist eine Standard-Bitcoin-Transaction, bereit zum Broadcast. Die PSBT hat ihren Zweck erfüllt.
