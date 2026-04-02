## Schnorr: Uma Assinatura Mais Simples

Assinaturas Schnorr (BIP340) foram ativadas com a atualizacao Taproot em novembro de 2021. Sao usadas para transacoes segwit v1 (P2TR).

O algoritmo Schnorr e matematicamente mais simples que o ECDSA. Ele produz uma assinatura fixa de 64 bytes (comparado aos 71-73 bytes variaveis do ECDSA). A equacao de verificacao e linear, o que permite agregacao de assinaturas: multiplas assinaturas podem ser combinadas em uma unica assinatura do mesmo tamanho.

Assinaturas Schnorr tem uma reducao de seguranca comprovavel para o problema do logaritmo discreto, uma garantia de seguranca mais forte que a do ECDSA.
