## Pourquoi les portefeuilles materiels en avaient besoin

Les portefeuilles materiels ont une contrainte specifique : ils ne peuvent pas interroger la blockchain. Ils recoivent des donnees, verifient ce qu'ils peuvent, signent et renvoient le resultat. Sans format standard, chaque fabricant de portefeuille materiel devait ecrire des integrations personnalisees pour chaque portefeuille logiciel. Ajouter un nouveau portefeuille materiel signifiait corriger chaque coordinateur. Ajouter un nouveau coordinateur signifiait corriger chaque portefeuille materiel. L'explosion combinatoire etait insoutenable.

PSBT a resolu cela en definissant un seul format qui transporte toutes les informations dont un signataire a besoin. L'UTXO depense, le chemin de derivation de la cle, le type de sighash a utiliser, le script de rachat pour P2SH, le script temoin pour P2WSH. Un portefeuille materiel recoit un PSBT, lit les champs necessaires, signe, ecrit sa signature partielle dans le PSBT et le renvoie. Pas de protocole proprietaire. Pas de negociation de format.

L'ecosysteme a converge rapidement. Coldcard, Trezor, Ledger, BitBox et Jade parlent tous PSBT. Les coordinateurs logiciels comme Sparrow, Specter et Bitcoin Core produisent et consomment tous des PSBT. Un quorum multisig peut utiliser du materiel de differents fabricants sans souci de compatibilite.
