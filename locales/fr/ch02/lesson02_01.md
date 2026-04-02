## Alice paie Bob

Alice veut acheter un produit dans la boutique en ligne de Bob. La page de paiement de Bob affiche une option de paiement Bitcoin avec un montant et une adresse Bitcoin.

Alice ouvre son application de portefeuille. Elle scanne la blockchain a la recherche de sorties de transaction non depensees (UTXO) verrouilees a ses adresses. Son portefeuille trouve un UTXO de 0.015 BTC, plus que suffisant pour couvrir le prix de Bob de 0.01 BTC.

Le portefeuille d'Alice construit une transaction avec une entree (referencant son UTXO de 0.015 BTC) et deux sorties : 0.01 BTC vers l'adresse de Bob, et 0.0049 BTC en retour a Alice comme monnaie. Les 0.0001 BTC restants constituent les frais de transaction.
