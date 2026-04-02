## Risques de consensus

Les applications qui valident des transactions doivent implementer les regles de consensus de maniere exacte. Une difference subtile dans la facon dont deux implementations gerent un cas limite peut les amener a diverger sur la validite d'un bloc, divisant le reseau de leur perspective.

L'approche la plus sure pour les applications : deleguer la validation de consensus a un noeud complet (Bitcoin Core ou equivalent) et utiliser son API pour les requetes blockchain. Ne reimplementez pas les regles de consensus dans le code applicatif a moins d'etre pret a correspondre a chaque cas limite de l'implementation de reference.

Pour les developpeurs de portefeuilles : utilisez des bibliotheques bien testees pour la derivation de cles, la generation d'adresses et la signature de transactions. Preferez les implementations open source eprouvees au code personnalise.
