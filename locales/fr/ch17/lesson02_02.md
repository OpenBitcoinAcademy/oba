## Interpolation de Lagrange

FROST s'appuie sur le partage de secret de Shamir. Pendant la generation de cles, la cle secrete du groupe est divisee en n parts de telle sorte que n'importe quelles t parts peuvent la reconstruire. L'outil mathematique qui rend cela possible est l'interpolation de Lagrange.

Un polynome de degre t-1 est uniquement determine par t points. Chaque signataire detient un point (sa part secrete) sur un polynome dont le terme constant est le secret du groupe. Pendant la signature, les signataires ne reconstruisent jamais le polynome complet. A la place, chaque signataire multiplie sa part de signature par un coefficient de Lagrange qui depend de l'ensemble des signataires participants.

Le coefficient de Lagrange pour le signataire $i$ dans un ensemble de signature $S$ est : $\lambda_i = \prod_{j \in S, j \neq i} \frac{j}{j - i}$. Ces coefficients assurent que la somme des parts de signature ponderees produit une signature Schnorr valide pour la cle publique du groupe, sans qu'aucune partie ne detienne ou reconstruise la cle secrete complete.
