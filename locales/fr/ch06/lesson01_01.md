## A quoi ressemble une transaction

Une transaction Bitcoin est une structure de donnees qui transfere des bitcoins d'un proprietaire a un autre. Elle contient des entrees (qui referencent des sorties precedentes a depenser), des sorties (qui creent de nouveaux montants depensables) et des donnees d'autorisation (des signatures prouvant que le depenseur controle les cles).

Les entrees indiquent d'ou viennent les bitcoins. Chaque entree reference la sortie d'une transaction precedente qui n'a pas encore ete depensee. Cette sortie non depensee est appelee un UTXO (unspent transaction output).

Les sorties indiquent ou vont les bitcoins. Chaque sortie specifie un montant en satoshis et une condition de verrouillage (un script) qui determine qui peut la depenser.

Une transaction consomme d'anciennes sorties et en cree de nouvelles. Rien n'est "stocke dans un compte". Bitcoin suit la propriete a travers une chaine de sorties, chacune verrouillee a une cle specifique.
