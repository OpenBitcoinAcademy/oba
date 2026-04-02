## Transactions d'engagement

Chaque partie detient sa propre version de la derniere transaction d'engagement. Ces transactions repartissent les fonds du canal selon le solde actuel. Si Alice a 0.7 BTC et Bob a 0.3 BTC, les deux transactions d'engagement refletent cette repartition.

Les transactions d'engagement sont asymetriques. La version d'Alice paie Bob immediatement mais verrouille les fonds d'Alice derriere un verrou temporel. La version de Bob fait l'inverse. Cette asymetrie permet le mecanisme de penalite : si Alice diffuse un ancien engagement (essayant de reclamer plus que ce qu'elle possede), Bob peut utiliser un secret de revocation pour prendre tous les fonds avant que le verrou temporel d'Alice n'expire.

A chaque mise a jour du solde, les deux parties echangent des secrets de revocation pour les anciens engagements, les rendant dangereux a diffuser.
