## Ce que sont les descripteurs de sortie

Un descripteur de sortie est une chaine qui decrit completement comment deriver un ensemble d'adresses Bitcoin. Les BIP 380 a 389 definissent le langage des descripteurs. Chaque descripteur specifie le type de script, les cles impliquees et les chemins de derivation.

Un descripteur comme `wpkh([d34db33f/84h/0h/0h]xpub.../0/*)` dit a un portefeuille tout ce qu'il doit savoir : utiliser P2WPKH, deriver de cette cle publique etendue a ce chemin, et generer des adresses sequentielles. Le portefeuille n'a pas besoin de deviner le type de script ou le schema de derivation. Le descripteur est autonome.

Avant les descripteurs, les portefeuilles s'appuyaient sur des conventions. BIP 44 disait que les portefeuilles HD devaient utiliser un chemin de derivation specifique. BIP 49 a ajoute un chemin different pour P2SH-SegWit. BIP 84 en a ajoute un autre pour le SegWit natif. Un portefeuille important un xpub devait essayer toutes les conventions et esperer avoir devine correctement. Les descripteurs remplacent les conjectures par des declarations explicites.
