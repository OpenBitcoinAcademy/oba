## P2PKH etape par etape

Voici ce qui se passe quand un script P2PKH s'execute :

1. La signature et la cle publique sont empilees sur la pile (depuis le script de deverrouillage).
2. OP_DUP duplique la cle publique sur la pile.
3. OP_HASH160 hache l'element du sommet (la cle publique dupliquee) avec SHA-256 puis RIPEMD-160.
4. Le hachage de cle publique attendu est empile sur la pile (depuis le script de verrouillage).
5. OP_EQUALVERIFY depile les deux elements du sommet et verifie qu'ils sont egaux. S'ils different, le script echoue.
6. OP_CHECKSIG depile la signature et la cle publique, verifie la signature par rapport aux donnees de la transaction. Si elle est valide, il empile vrai.

Le resultat : vrai sur la pile. La depense est autorisee. La signature prouve que le depenseur controle la cle privee, sans la reveler.
