## Was ist ein Block?

Ein Block ist ein Bündel von Transactions, die zusammengefasst und der Blockchain hinzugefügt werden. Jeder Block enthält einen Header und eine Liste von Transactions.

Der Block-Header hat sechs Felder: Version, vorheriger Block-Hash, Merkle Root (ein Hash aller Transactions im Block), Zeitstempel, Difficulty-Target und Nonce. Der vorherige Block-Hash verknüpft jeden Block mit seinem Vorgänger und bildet so eine Kette.

Miner konkurrieren darum, einen Nonce-Wert zu finden, der den Hash des Block-Headers unter das Difficulty-Target fallen lässt. Das erfordert Billionen von Versuchen. Der erste Miner, der eine gültige Nonce findet, sendet den Block. Andere Nodes verifizieren ihn und fügen ihn ihrer Kopie der Chain hinzu.
