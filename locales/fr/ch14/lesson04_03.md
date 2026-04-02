## Mises a jour futures via OP_SUCCESS

Tapscript reserve un ensemble d'opcodes appeles OP_SUCCESSx. Si un script contient un opcode OP_SUCCESS, le script entier reussit immediatement sans validation supplementaire.

C'est un mecanisme de mise a jour delibere. Un futur soft fork peut redefinir un opcode OP_SUCCESSx pour effectuer une nouvelle validation (comme OP_CHECKTEMPLATEVERIFY ou OP_VAULT). Les anciens noeuds voient OP_SUCCESS et acceptent la transaction. Les nouveaux noeuds executent l'opcode redefini et appliquent les nouvelles regles.

Taproot supporte aussi les versions de feuille inconnues. Le Tapscript actuel utilise la version de feuille 0xC0. Un futur soft fork peut definir la version de feuille 0xC2 avec une semantique de script entierement differente. Les anciens noeuds ignorent la validation des versions de feuille inconnues. Les nouveaux noeuds appliquent les nouvelles regles.

Les deux mecanismes permettent a Taproot d'evoluer sans le remplacer. Chaque version de feuille de script et chaque opcode OP_SUCCESS est un emplacement reserve pour une future fonctionnalite.
