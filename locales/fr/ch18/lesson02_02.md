## Propagation de la preimage

Carol connait R. Elle le revele a Bob et reclame 1.000 BTC. Bob connait maintenant R. Il le revele a Alice et reclame 1.001 BTC. La preimage se propage en arriere le long du chemin de paiement.

Alice a paye 1.001 BTC. Carol a recu 1.000 BTC. Bob a gagne 0.001 BTC comme frais de routage. Le paiement s'est regle en quelques secondes sans toucher la blockchain.

Tous les HTLC de la chaine utilisent le meme hachage H(R). Soit la preimage se propage jusqu'au bout (le paiement reussit entierement), soit les verrous temporels expirent (le paiement echoue entierement, tous les fonds sont rendus). Aucun etat intermediaire n'est possible. Le paiement est atomique.
