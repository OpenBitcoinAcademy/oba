## Vie privee par divulgation selective

Une sortie Taproot avec 8 feuilles de script ne revele qu'une seule feuille lors de la depense par chemin de script. Les 7 autres feuilles apparaissent comme des hachages dans la preuve de Merkle. Un observateur apprend une condition de depense mais ne peut pas determiner quelles etaient les autres conditions.

Comparez avec le multisig P2WSH : la depense revele le script complet, toutes les cles publiques et quelles cles ont signe. Chaque partie impliquee est visible.

Un multisig 2-of-3 utilisant Taproot : le chemin de cle contient l'agregat MuSig2 des deux signataires les plus courants. Deux feuilles de script contiennent les paires de repli (A+C et B+C). Dans le cas courant, le chemin de cle est utilise et rien du multisig n'est revele. En cas de repli, seule une paire alternative est exposee. L'autre paire reste un hachage.
