## Le mempool

Quand vous diffusez une transaction, elle n'entre pas dans un bloc immediatement. Elle entre dans le mempool : une zone d'attente ou les transactions non confirmees attendent qu'un mineur les inclue dans un bloc.

Chaque noeud maintient son propre mempool. Les transactions se propagent a travers le reseau en quelques secondes, mais la confirmation (inclusion dans un bloc) prend en moyenne 10 minutes. Pendant les periodes chargees, le mempool grossit et les transactions a frais bas peuvent attendre plus longtemps.

Une transaction dans le mempool est non confirmee. Elle a ete validee par les noeuds (signatures correctes, entrees non depensees) mais n'a pas encore ete ecrite dans la blockchain.
