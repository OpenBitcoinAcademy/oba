## Le systeme de types

Chaque fragment Miniscript a un type qui decrit son comportement sur la pile Script. Les quatre types de base sont B, V, K et W.

Une expression de type B (base) empile une seule valeur non nulle en cas de succes et zero en cas d'echec. Une expression de type V (verify) reussit sans sortie sur la pile ou interrompt le script. Une expression de type K (key) empile une cle publique en cas de succes. Une expression de type W (wrapped) se comporte comme B mais consomme d'abord l'element du sommet de la pile, utilisee pour combiner des sous-expressions.

Le compilateur verifie les types a chaque point de composition. Un fragment `and_v` exige que son premier argument soit de type V et son second de type B. Si vous passez des arguments de mauvais types, le compilateur rejette l'expression. Cela detecte des erreurs qui, en Script brut, produiraient un script qui "fonctionne" mais applique les mauvaises conditions.

Le systeme de types suit aussi des proprietes comme la dissipabilite (une branche non satisfaite peut-elle laisser la pile propre ?) et la non-malleabilite (y a-t-il exactement un temoin valide pour chaque chemin de depense ?). Ces proprietes se propagent a travers la composition. Une expression Miniscript les possede ou non, et le compilateur peut rapporter lesquelles.
