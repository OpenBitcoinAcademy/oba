## Trois couches : politique, Miniscript, Script

Miniscript fonctionne dans une architecture a trois couches. Au sommet se trouve un langage de politique lisible par l'humain. Au milieu se trouve Miniscript lui-meme. En bas se trouve Bitcoin Script.

Une politique decrit ce que vous voulez : "Alice ET Bob, OU Carol apres 90 jours." Vous l'ecrivez comme `or(and(pk(Alice),pk(Bob)),and(pk(Carol),older(12960)))`. Le langage de politique est pour les humains. Il nomme les cles et les verrous temporels sans se soucier des opcodes.

Un compilateur traduit la politique en une expression Miniscript : `or_d(and_v(v:pk(Alice),pk(Bob)),and_v(v:pk(Carol),older(12960)))`. L'expression Miniscript specifie exactement comment les conditions sont composees, y compris quels types de fragments sont utilises a chaque position. A partir de la, l'expression se traduit directement en opcodes Bitcoin Script. Chaque fragment Miniscript a un et un seul encodage Script.
