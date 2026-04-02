## Un nouveau langage de script

Tapscript (BIP 342) definit les regles de validation pour les scripts executes dans les depenses par chemin de script de Taproot. Il est similaire au Bitcoin Script legacy mais avec des ameliorations ciblees.

Tous les opcodes de verification de signature (OP_CHECKSIG, OP_CHECKSIGVERIFY) utilisent les signatures Schnorr au lieu d'ECDSA. Un nouvel opcode, OP_CHECKSIGADD, remplace le legacy OP_CHECKMULTISIG. Au lieu de verifier plusieurs signatures contre plusieurs cles en une seule operation, OP_CHECKSIGADD verifie une signature a la fois et accumule un compteur. Cela permet la verification par lots : le verificateur collecte toutes les signatures et les verifie ensemble en une seule operation, plus rapidement que de verifier chacune individuellement.
