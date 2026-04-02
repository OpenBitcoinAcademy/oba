## Replace By Fee (RBF)

Si une transaction est bloquee dans le mempool parce que son taux de frais est trop bas, l'expediteur peut la remplacer par une nouvelle version qui paie des frais plus eleves. C'est le Replace By Fee (RBF), defini dans BIP125.

Pour utiliser RBF, la transaction originale doit signaler la remplacabilite en reglant le numero de sequence d'au moins une entree a une valeur inferieure a 0xFFFFFFFE. Le logiciel de portefeuille gere cela automatiquement.

La transaction de remplacement doit payer des frais totaux plus eleves que l'originale. Elle peut modifier les sorties (payer un montant different ou ajouter/retirer des destinataires) tant qu'elle depense au moins une des memes entrees.
