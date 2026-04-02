## Le flux de travail en lecture seule

Un flux de travail de signature par portefeuille materiel commence avec un portefeuille en lecture seule sur un ordinateur connecte au reseau. Le portefeuille en lecture seule detient les cles publiques etendues (xpubs) de tous les signataires mais aucune cle privee. Il peut generer des adresses, suivre les soldes et construire des transactions, mais il ne peut pas signer.

Quand l'utilisateur veut depenser, le portefeuille en lecture seule cree un PSBT. Il remplit la transaction non signee, attache les donnees WITNESS_UTXO pour chaque entree et ecrit les entrees BIP32_DERIVATION pour que le portefeuille materiel sache quelles cles utiliser. Le PSBT est exporte vers un fichier sur carte SD ou encode comme un code QR anime.

Le portefeuille materiel recoit le PSBT, analyse les entrees et sorties, et affiche un resume a l'utilisateur : quelles adresses recoivent des fonds, combien va a chacune et combien est paye en frais. L'utilisateur confirme sur l'appareil. Le portefeuille materiel derive les cles privees de sa graine en utilisant les chemins BIP 32 du PSBT, signe chaque entree qu'il peut, ecrit les entrees PARTIAL_SIG et exporte le PSBT mis a jour via carte SD ou code QR.

Le portefeuille en lecture seule importe le PSBT signe. Si toutes les signatures requises sont presentes, il finalise et extrait la transaction brute, puis la diffuse. Si plus de signatures sont necessaires (comme dans une configuration multisig), il transmet le PSBT au signataire suivant.
