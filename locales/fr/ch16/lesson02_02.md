## Metteur a jour et Signataire

Le Metteur a jour ajoute les metadonnees dont les Signataires ont besoin. Pour chaque entree, le Metteur a jour attache l'UTXO depense (soit la transaction precedente complete, soit l'UTXO temoin specifique), le chemin de derivation BIP 32 de la cle de signature, le script de rachat si l'entree est P2SH, et le script temoin si l'entree est P2WSH. Pour chaque sortie, le Metteur a jour peut attacher des chemins de derivation BIP 32 pour que le signataire puisse verifier que les sorties de monnaie appartiennent au meme portefeuille.

Le Signataire lit le PSBT, identifie les entrees qu'il peut signer et produit des signatures partielles. Pour chaque entree qu'il signe, il ecrit une entree PARTIAL_SIG indexee par la cle publique. Le Signataire ne modifie pas la transaction elle-meme. Il ne finalise pas les scripts. Il ajoute sa signature et transmet le PSBT au participant suivant.

Un Signataire doit verifier les donnees UTXO avant de signer. Si le PSBT pretend qu'une entree depense 0.5 BTC mais que l'UTXO reel contient 1.0 BTC, le signataire ferait involontairement don de 0.5 BTC en frais. Les portefeuilles materiels verifient les montants UTXO par rapport aux montants des sorties et avertissent l'utilisateur si les frais semblent deraisonnables.
