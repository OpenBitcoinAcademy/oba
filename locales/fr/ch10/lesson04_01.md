## Clients SPV

Tous les appareils ne peuvent pas stocker et valider la blockchain complete. Les telephones mobiles, les appareils embarques et les portefeuilles materiels ont un stockage et une puissance de traitement limites. Ces appareils executent des clients **Simplified Payment Verification** (SPV).

Un client SPV telecharge uniquement les en-tetes de blocs, pas les blocs complets. La chaine d'en-tetes est petite : environ 60 Mo pour l'historique complet de Bitcoin. Le client peut verifier qu'un en-tete de bloc satisfait la cible de preuve de travail, confirmant qu'un mineur a depense de l'energie reelle pour le produire.

Pour verifier si une transaction est confirmee, le client SPV demande une **preuve de Merkle** a un noeud complet. La preuve montre que le hachage de la transaction est inclus dans la racine de Merkle d'un bloc. Le client fait confiance au fait que les noeuds complets ont valide les transactions du bloc, car falsifier un en-tete de preuve de travail valide est prohibitivement couteux.

SPV donne aux clients legers un niveau raisonnable d'assurance sans telecharger la blockchain complete. Le compromis est qu'un client SPV ne peut pas detecter si un bloc contient une transaction invalide. Il fait confiance au poids economique de la preuve de travail plutot que de verifier chaque regle lui-meme.
