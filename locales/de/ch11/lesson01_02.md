## Der Mempool

Wenn du eine Transaction sendest, gelangt sie nicht sofort in einen Block. Sie tritt in den Mempool ein: einen Wartebereich, in dem unbestätigte Transactions liegen, bis ein Miner sie in einen Block aufnimmt.

Jede Node pflegt ihren eigenen Mempool. Transactions verbreiten sich innerhalb von Sekunden im Netzwerk, aber die Bestätigung (Aufnahme in einen Block) dauert durchschnittlich 10 Minuten. In geschäftigen Zeiten wächst der Mempool und Transactions mit niedrigen Fees warten möglicherweise länger.

Eine Transaction im Mempool ist unbestätigt. Sie wurde von Nodes validiert (korrekte Signatures, nicht ausgegebene Inputs), aber noch nicht in die Blockchain geschrieben.
