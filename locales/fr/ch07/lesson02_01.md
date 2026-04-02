## Pay to Script Hash (P2SH)

P2SH separe le script de l'adresse. Au lieu d'encoder le script de verrouillage complet dans la sortie, P2SH encode uniquement un hachage du script. Le depenseur revele le script complet (appele script de rachat) lors de la depense.

Le script de sortie contient : OP_HASH160 <script_hash> OP_EQUAL. Pour depenser, l'entree fournit le script de rachat et les donnees requises par ce script (signatures, cles publiques). Le reseau hache le script de rachat fourni et verifie qu'il correspond au hachage dans la sortie.

Les adresses P2SH commencent par "3" sur le mainnet. Elles permettent des conditions de depense complexes (multisig, timelocks) sans que l'expediteur ait besoin de connaitre les details. L'expediteur paie a un hachage court. Le destinataire gere la complexite.
