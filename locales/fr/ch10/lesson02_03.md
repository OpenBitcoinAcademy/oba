## Synchronisation de la blockchain

Un nouveau noeud doit telecharger et valider l'integralite de la blockchain avant de pouvoir verifier les transactions actuelles. Ce processus est appele **Initial Block Download** (IBD).

Le noeud envoie des messages `getheaders` a ses pairs, demandant des en-tetes de blocs par lots. Les en-tetes sont petits (80 octets chacun) et arrivent rapidement. Le noeud les utilise pour construire la chaine d'en-tetes avec le plus de preuve de travail cumulee.

Une fois que le noeud identifie la meilleure chaine d'en-tetes, il demande les blocs complets en utilisant des messages `getdata`. Il telecharge des blocs depuis plusieurs pairs en parallele pour accelerer le processus. Chaque bloc est valide par rapport aux regles de consensus a son arrivee : signatures de transaction, execution de scripts, limites de poids et cible de preuve de travail.

L'IBD peut prendre des heures ou des jours sur du materiel lent. Une fois termine, le noeud passe en fonctionnement normal. Il recoit les nouveaux blocs et transactions au fur et a mesure de leur diffusion et les valide en temps reel.
