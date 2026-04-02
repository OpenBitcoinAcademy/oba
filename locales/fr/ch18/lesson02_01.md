## Paiements conditionnels a travers les sauts

Les Hash Time-Locked Contracts (HTLC) permettent des paiements entre des parties qui ne partagent pas de canal direct. Un HTLC combine deux conditions : un hashlock (reveler un secret pour reclamer les fonds) et un timelock (si le secret n'est pas revele a temps, les fonds retournent a l'expediteur).

Carol genere un secret aleatoire R et envoie le hachage H(R) a Alice dans une facture de paiement. Alice ne connait pas R. Elle cree un HTLC avec Bob : "Paie 1.001 BTC si tu reveles la preimage de H(R) dans les 40 blocs." Bob cree un HTLC avec Carol : "Paie 1.000 BTC si tu reveles la preimage de H(R) dans les 30 blocs."
