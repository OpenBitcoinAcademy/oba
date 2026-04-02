## Fee-Schätzung

Wallet-Software schätzt Fee-Raten, indem sie den Mempool analysiert: die Menge der unbestätigten Transactions, die auf die Aufnahme in einen Block warten.

Ein gängiger Ansatz: die Fee-Raten der kürzlich in Blocks aufgenommenen Transactions und die Fee-Raten der noch wartenden Transactions betrachten. Dann eine Fee-Rate wählen, die hoch genug ist, um mit der niedrigsten Fee-Rate-Transaction im letzten Block zu konkurrieren, mit einem Sicherheitspuffer.

Fee-Schätzung ist nicht perfekt. Die Nachfrage schwankt. Ein Nutzungsanstieg kann eine zuvor ausreichende Fee-Rate unzureichend machen. Manche Wallets erlauben es, eigene Fee-Raten festzulegen.
