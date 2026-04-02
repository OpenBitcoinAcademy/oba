## Cashu : ecash a atelier unique

Cashu est un protocole d'ecash Chaumien utilisant un seul operateur d'atelier au lieu d'une federation. Les utilisateurs deposent des bitcoins (generalement via Lightning) et recoivent des jetons ecash signes a l'aveugle. Les jetons sont des instruments au porteur : quiconque detient un jeton valide peut le racheter.

L'atelier ne peut pas lier la frappe au rachat (les signatures aveugles le garantissent). L'atelier ne peut pas voir qui transfere des jetons a qui. Les transferts entre utilisateurs sont instantanes et gratuits (pas d'interaction avec la blockchain).

Le compromis : les utilisateurs font confiance a un seul operateur d'atelier. Si l'atelier est malhonnete ou se deconnecte, les fonds peuvent etre perdus. Cashu est concu pour des deployments a petite echelle et specifiques a une application (un cafe, un service de streaming, une communaute) ou la relation de confiance est gerable.
