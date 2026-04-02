## Les entrees referencent des sorties precedentes

Chaque entree pointe vers une sortie specifique d'une transaction precedente. Elle le fait avec deux informations : l'identifiant de transaction (un hachage) et l'index de sortie (quelle sortie dans cette transaction).

Quand vous depensez des bitcoins, vous prouvez que vous controlez la cle capable de deverrouiller une sortie precedente. Dans les transactions legacy, la preuve (signature et cle publique) se trouve dans le script d'entree. Dans les transactions segwit modernes, la preuve se trouve dans une structure temoin separee, et le script d'entree est vide.

Une fois qu'une sortie est referencee par une entree, elle est depensee. Elle ne peut pas etre depensee a nouveau. C'est ainsi que Bitcoin empeche la double depense sans autorite centrale.
