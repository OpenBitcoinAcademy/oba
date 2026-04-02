## La transaction coinbase

Chaque bloc commence par une transaction speciale appelee la transaction coinbase. Elle n'a pas d'entrees provenant de transactions precedentes. A la place, elle cree de nouveaux bitcoins comme recompense pour le mineur qui a trouve le bloc.

La transaction coinbase a une entree avec une reference nulle (pas de sortie precedente depensee). Le mineur remplit le champ de script de cette entree avec des donnees arbitraires, incluant l'extra nonce utilise pendant le minage. Satoshi Nakamoto a fameusement integre un titre de journal dans le coinbase du bloc de genese : "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."

Les sorties de la transaction coinbase paient le mineur. La valeur totale des sorties ne doit pas depasser la subvention de bloc plus la somme de tous les frais de transaction du bloc. Toute valeur non reclamee est detruite.
