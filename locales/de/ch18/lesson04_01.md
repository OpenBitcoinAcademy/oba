## Einen Channel öffnen

Das Öffnen eines Channels erfordert eine On-Chain-Transaction: die Funding Transaction. Sie erzeugt einen 2-of-2-Multisig-Output. Vor dem Broadcast signieren beide Parteien die erste Commitment Transaction (das Sicherheitsnetz, das Mittel zurückgibt, falls der Channel scheitert).

Der Channel ist offen, sobald die Funding Transaction bestätigt ist. Beide Parteien können das Guthaben aktualisieren, indem sie neue Commitment Transactions austauschen. Die On-Chain-Kosten betragen eine Transaktionsgebühr für die Funding Transaction.
