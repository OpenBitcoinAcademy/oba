## Package Relay

Bisher hat Bitcoin Core jede Transaction einzeln für die Mempool-Aufnahme bewertet. Eine Parent-Transaction mit niedriger Fee wurde abgelehnt, selbst wenn eine Child-Transaction mit hoher Fee von ihr abhing.

Package Relay ändert das. Nodes bewerten Gruppen zusammengehöriger Transactions gemeinsam und akzeptieren einen Parent mit niedriger Fee, wenn das Child die kombinierte Fee-Rate über den Schwellenwert bringt.

Das verbessert die Zuverlässigkeit von CPFP. Ohne Package Relay könnte der Parent die Miner nicht erreichen, was das Child nutzlos macht. Mit Package Relay wandern Parent und Child gemeinsam durch das Netzwerk.
