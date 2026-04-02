## Privacidade por Revelacao Seletiva

Um output Taproot com 8 folhas de script revela apenas uma folha quando gasto via script path. As outras 7 folhas aparecem como hashes na prova Merkle. Um observador descobre uma condicao de gasto, mas nao pode determinar quais eram as outras condicoes.

Compare com multisig P2WSH: gastar revela o script completo, todas as chaves publicas e quais chaves assinaram. Toda parte envolvida e visivel.

Um multisig 2-de-3 usando Taproot: o key path guarda o agregado MuSig2 dos dois assinantes mais comuns. Duas folhas de script guardam os pares alternativos (A+C e B+C). No caso comum, o key path e usado e nada sobre o multisig e revelado. Em uma alternativa, apenas um par alternativo e exposto. O outro par permanece como hash.
