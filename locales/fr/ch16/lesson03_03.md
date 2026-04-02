## Champs Taproot et differences de version

BIP 371 a ajoute des champs specifiques a Taproot au format PSBT. TAP_KEY_SIG (type de cle d'entree 0x13) stocke une signature Schnorr pour une depense par chemin de cle. TAP_SCRIPT_SIG (type de cle d'entree 0x14) stocke une signature Schnorr pour une feuille de script specifique, indexee par la cle publique X-only et le hachage de feuille.

TAP_LEAF_SCRIPT (type de cle d'entree 0x15) fournit le script, sa version de feuille et le bloc de controle necessaire pour la depense par chemin de script. TAP_BIP32_DERIVATION (type de cle d'entree 0x16) etend le champ de derivation BIP 32 standard avec une liste de hachages de feuilles pour lesquels la cle peut signer. TAP_INTERNAL_KEY (type de cle d'entree 0x17) enregistre la cle publique interne avant modification.

Cote sortie, TAP_INTERNAL_KEY (type de cle de sortie 0x05) et TAP_BIP32_DERIVATION (type de cle de sortie 0x07) permettent aux signataires de verifier que les sorties de monnaie Taproot appartiennent au meme portefeuille. Le signataire peut re-deriver la cle modifiee a partir de la cle interne et confirmer qu'elle correspond au scriptPubKey de la sortie.

PSBTv2 (BIP 370) differe de v0 dans sa structure, pas dans son concept. Dans v0, la transaction non signee se trouve dans la carte globale sous forme d'un seul blob serialise. Dans v2, les entrees et sorties sont decrites par des champs par carte : PREVIOUS_TXID, OUTPUT_INDEX, SEQUENCE pour les entrees ; AMOUNT et SCRIPT pour les sorties. Les Constructeurs peuvent ajouter des entrees et sorties de maniere incrementale sans re-serialiser la transaction entiere a chaque fois.
