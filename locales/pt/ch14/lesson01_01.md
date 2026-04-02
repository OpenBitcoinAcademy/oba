## A Chave Interna e o Tweak

Um output Taproot bloqueia fundos em uma chave publica ajustada Q. Essa chave e derivada de duas entradas: uma chave publica interna P e uma raiz Merkle opcional de uma arvore de scripts.

O valor do tweak t e calculado como: t = tagged_hash("TapTweak", P || merkle_root). A chave ajustada e: Q = P + t * G, onde G e o ponto gerador na secp256k1.

Se nao ha arvore de scripts, o tweak usa apenas P: t = tagged_hash("TapTweak", P). O output ainda se compromete com a chave interna, mas nenhum script esta incorporado.

Na cadeia, o scriptPubKey e: OP_1 seguido pela coordenada X de 32 bytes de Q. Sao 34 bytes, o mesmo tamanho de um output P2WSH. Nenhum observador pode dizer se Q contem uma arvore de scripts incorporada ou nao.
