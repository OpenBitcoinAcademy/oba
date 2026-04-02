## Creation et verification

Creer une signature necessite deux entrees : la cle privee et le message (les donnees de la transaction, ou un hachage de celles-ci). L'algorithme de signature produit une valeur de signature unique a la fois pour la cle et le message.

Verifier une signature necessite trois entrees : la cle publique, le message et la signature. L'algorithme de verification renvoie vrai ou faux. Si vrai, la signature a ete produite par le detenteur de la cle privee correspondante pour ce message exact.

La cle privee n'est jamais revelee lors de la creation ou de la verification. N'importe qui peut verifier, mais seul le detenteur de la cle peut signer.
