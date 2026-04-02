## Politiques de depense a decroissance temporelle

Une politique de recuperation a verrou temporel commence avec une condition de depense principale et ajoute des conditions de repli qui s'activent apres un delai. La cle principale controle les fonds immediatement. Si la cle principale est perdue ou que son detenteur devient indisponible, un chemin de depense alternatif se deverrouille apres un nombre specifie de blocs.

En Miniscript, cela s'exprime comme : `or(pk(primary),and(pk(recovery),older(26280)))`. La cle principale peut depenser a tout moment. La cle de recuperation ne peut depenser qu'apres environ six mois (26 280 blocs a 10 minutes par bloc). La politique se compile en un seul Script avec deux chemins d'execution.

Ce modele resout un vrai probleme. Un portefeuille a cle unique n'a pas de chemin de recuperation. Si la cle est perdue, les fonds sont partis. Un multisig standard necessite plusieurs signataires pour chaque transaction, ajoutant de la friction a l'utilisation quotidienne. Un repli a verrou temporel donne au detenteur principal le controle total en fonctionnement normal tout en garantissant qu'une partie de confiance peut recuperer les fonds apres un delai.
