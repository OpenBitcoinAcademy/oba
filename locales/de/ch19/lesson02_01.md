## Single-Use Seals

RGB definiert ein "Seal" als einen Bitcoin-UTXO. Wenn ein Zustandsübergang stattfindet (ein Token wird übertragen, ein Vertrag wird aktualisiert), wird das Seal "geschlossen", indem dieser UTXO ausgegeben wird. Ein UTXO kann nur einmal ausgegeben werden, also kann ein Seal nur einmal geschlossen werden. Das verhindert Double-Spending von RGB-Assets mithilfe von Bitcoins bestehendem Konsensmechanismus.

Die Zustandsübergangsdaten (wem gehört was, Vertragsaktualisierungen) berühren nie die Blockchain. Nur ein kryptographisches Commitment auf den Übergang wird in die Bitcoin-Transaction eingebettet, typischerweise in einen Taproot-Output. Die Blockchain verankert den Zeitpunkt und die Reihenfolge. Die Daten liegen off-chain bei den beteiligten Parteien.
