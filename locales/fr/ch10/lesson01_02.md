## Types de noeuds

Tous les noeuds n'executent pas le meme logiciel et ne remplissent pas le meme role. Les noeuds different par les donnees qu'ils stockent et les fonctions qu'ils fournissent.

Un **noeud complet** telecharge et valide chaque bloc et chaque transaction par rapport aux regles de consensus. Il ne fait confiance a personne. Il peut verifier independamment n'importe quel paiement. Bitcoin Core est l'implementation de noeud complet la plus courante.

Un **noeud mineur** est un noeud complet qui rivalise aussi pour creer de nouveaux blocs. Il assemble des blocs candidats a partir du mempool et effectue la preuve de travail. Les noeuds mineurs gagnent la subvention de bloc et les frais de transaction quand ils trouvent un bloc valide.

Un **noeud leger** (aussi appele client SPV) ne telecharge pas les blocs complets. Il telecharge uniquement les en-tetes de blocs et demande la preuve que des transactions specifiques existent. Les noeuds legers font confiance aux noeuds complets dans une certaine mesure, echangeant la securite contre une utilisation moindre des ressources.
