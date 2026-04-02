## Pay to Script Hash (P2SH)

P2SH trennt das Script von der Address. Statt das vollständige Locking Script im Output zu kodieren, kodiert P2SH nur einen Hash des Scripts. Der Absender enthüllt das vollständige Script (genannt Redeem Script) erst beim Ausgeben.

Das Output Script enthält: OP_HASH160 <script_hash> OP_EQUAL. Zum Ausgeben liefert der Input das Redeem Script und alle Daten, die das Redeem Script benötigt (Signatures, Public Keys). Das Netzwerk hasht das bereitgestellte Redeem Script und prüft, ob es mit dem Hash im Output übereinstimmt.

P2SH-Addresses beginnen mit "3" im Mainnet. Sie ermöglichen komplexe Ausgabebedingungen (Multisig, Timelocks), ohne dass der Sender die Details kennen muss. Der Sender zahlt an einen kurzen Hash. Der Empfänger handhabt die Komplexität.
