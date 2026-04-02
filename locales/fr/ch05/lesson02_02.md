## Generation deterministe de cles

La generation deterministe de cles resout le probleme de sauvegarde. Une graine aleatoire unique produit une sequence de cles privees par hachage repete. Chaque cle de la sequence peut etre recree a partir de la graine.

Sauvegardez la graine une seule fois et vous pouvez recuperer chaque cle que le portefeuille a generee ou generera. La graine est le secret maitre. Perdez-la et vous perdez l'acces a toutes les cles derivees.

La forme la plus elementaire : commencez avec une graine, hachez-la pour obtenir la cle 1, hachez la cle 1 pour obtenir la cle 2, et ainsi de suite. Cela produit une liste plate de cles. Mais les portefeuilles modernes utilisent une approche plus sophistiquee : un arbre.
