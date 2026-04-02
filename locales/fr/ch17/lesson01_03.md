## Vie privee par indiscernabilite

Le gain de vie privee des signatures a seuil va au-dela de la dissimulation de la politique de depense. Avec OP_CHECKMULTISIG, le modele de script lui-meme est une empreinte. Les entreprises d'analyse de chaine categorisent les adresses par type de script, identifiant les portefeuilles multisig et deduisant les arrangements de garde.

FROST elimine cette empreinte. Un portefeuille de garde 2-of-3, une tresorerie d'entreprise 5-of-7 et un portefeuille personnel a cle unique produisent tous des sorties identiques sur la chaine. Chacun depense avec une cle publique de 32 octets et une signature de 64 octets dans un chemin de cle Taproot.

Cette indiscernabilite beneficie a tous les utilisateurs de Taproot. Plus l'ensemble de transactions qui se ressemblent est grand, plus il est difficile de distinguer une transaction individuelle. Les utilisateurs de signatures a seuil se fondent dans la foule des depenseurs a cle unique, et les depenseurs a cle unique gagnent un deni plausible qu'ils pourraient etre des signataires a seuil.
