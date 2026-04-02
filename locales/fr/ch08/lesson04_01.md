## Types de SIGHASH

Une signature n'engage pas sur la transaction entiere par defaut. Le drapeau SIGHASH ajoute a chaque signature specifie quelles parties de la transaction la signature couvre.

**SIGHASH_ALL** (0x01) est la valeur par defaut. La signature engage sur toutes les entrees et toutes les sorties. Modifier n'importe quelle partie de la transaction invalide la signature.

**SIGHASH_NONE** (0x02) engage sur toutes les entrees mais aucune sortie. N'importe qui peut modifier les sorties apres la signature. Utilise dans de rares protocoles collaboratifs.

**SIGHASH_SINGLE** (0x03) engage sur toutes les entrees mais uniquement la sortie au meme index que l'entree signee. Les autres sorties peuvent etre modifiees.

Le modificateur ANYONECANPAY peut se combiner avec n'importe lequel de ceux-ci, permettant des signatures qui engagent sur une seule entree.
