## Ce que Miniscript resout

Miniscript, specifie dans BIP 379, est une representation structuree de Bitcoin Script. Il fait correspondre un ensemble defini de fragments composables aux opcodes de Script. Chaque fragment a des proprietes connues : son type, son cout en ressources et les donnees temoin requises.

Parce que la correspondance est structuree, un logiciel peut analyser une expression Miniscript et determiner chaque chemin de depense, chaque cle requise, chaque contrainte de verrou temporel et la taille exacte du temoin pour chaque chemin. Deux expressions qui encodent la meme politique peuvent etre comparees et montrees equivalentes.

La composition devient sure. Un portefeuille peut prendre deux fragments Miniscript, les combiner avec AND ou OR, et calculer la consommation de ressources du script resultant avant la diffusion. Si un chemin d'execution depasse les limites de consensus, le compilateur rejette la composition au moment de la construction, pas au moment de la depense.
