## Elargir le multisig avec le temps

Les politiques a verrou temporel peuvent elargir leur ensemble de signataires au fil du temps. Une entreprise pourrait commencer avec une politique 2-of-2 entre deux cofondateurs. Apres un an, une troisieme cle (detenue par un conseiller juridique) s'active, et la politique devient 2-of-3.

L'expression Miniscript est : `or(multi(2,founder_a,founder_b),and(multi(2,founder_a,founder_b,counsel),older(52560)))`. Pendant la premiere annee, seuls les deux fondateurs peuvent signer. Apres 52 560 blocs (environ un an), deux des trois cles suffisent pour signer.

La politique entiere vit dans un seul UTXO. Aucune transaction sur la chaine n'est necessaire quand le verrou temporel expire. Le chemin de depense supplementaire a ete engage au moment de la creation de la sortie. La cle du conseiller gagne le pouvoir de depense automatiquement quand la hauteur de bloc depasse le seuil. Cela elimine le besoin d'une ceremonie active de cles au point de transition.

Miniscript rend ces compositions auditables. Le compilateur peut enumerer chaque chemin de depense et ses conditions. Un auditeur peut verifier que la cle du conseiller n'a aucun pouvoir de depense avant le verrou temporel, que les fondateurs conservent le controle total en permanence, et que les tailles de temoin restent dans les limites de consensus.
