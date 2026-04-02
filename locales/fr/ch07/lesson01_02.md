## Scripts de deverrouillage

Pour depenser une sortie verrouillee, le depenseur fournit un script de deverrouillage, appele scriptSig. Pour P2PKH, le script de deverrouillage contient une signature numerique et la cle publique du depenseur.

Bitcoin execute ces scripts sur une machine a pile. D'abord, le script de deverrouillage s'execute et empile des donnees sur la pile. Puis la pile est copiee, et le script de verrouillage s'execute sur cette pile copiee. Les deux scripts ne se combinent jamais en un seul. Cette separation a ete introduite en 2010 pour corriger une vulnerabilite de securite.

Si le script de verrouillage se termine avec une valeur vraie au sommet de la pile, la depense est valide. Si la pile est vide ou si la valeur au sommet est zero, la depense echoue.
