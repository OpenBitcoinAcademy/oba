## Estimation des frais

Le logiciel de portefeuille estime les taux de frais en analysant le mempool : l'ensemble des transactions non confirmees en attente d'inclusion dans un bloc.

Une approche courante : examiner les taux de frais des transactions recemment incluses dans les blocs et les taux de frais des transactions encore en attente. Definir un taux de frais assez eleve pour rivaliser avec la transaction au taux le plus bas du bloc le plus recent, avec une marge de securite.

L'estimation des frais est imparfaite. La demande fluctue. Un pic d'utilisation peut rendre un taux de frais precedemment adequat insuffisant. Certains portefeuilles permettent aux utilisateurs de definir des taux de frais personnalises pour plus de controle.
