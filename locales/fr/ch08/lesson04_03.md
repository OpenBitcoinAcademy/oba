## L'algorithme de signature de SegWit

SegWit a introduit un nouvel algorithme de signature (BIP143) qui corrige un probleme de longue date : dans les transactions legacy, le montant depense n'est pas inclus dans les donnees signees. Cela forcait les signataires a recuperer la transaction precedente complete pour verifier le montant, ralentissant les portefeuilles materiels.

BIP143 inclut le montant de chaque entree dans les donnees signees. Un portefeuille materiel peut verifier la valeur totale des entrees et les frais sans telecharger les transactions precedentes. Le processus de signature est plus rapide et plus securise.

Pour segwit v1 (Taproot), BIP341 definit un algorithme mis a jour qui inclut des engagements supplementaires, supportant les depenses par chemin de cle et par chemin de script.
