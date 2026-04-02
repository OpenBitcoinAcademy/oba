## BIP32: O Padrao de Carteira HD

O BIP32 define como derivar uma arvore de chaves a partir de uma unica semente. O processo usa HMAC-SHA512 para dividir cada etapa de derivacao em uma chave filha e um chain code. O chain code e misturado na proxima derivacao, impedindo que qualquer pessoa compute chaves irmas.

A chave-mestra fica na raiz da arvore. A partir dela, uma sequencia de chaves filhas e derivada, cada uma identificada por um numero de indice (0, 1, 2, ...). Cada filha pode produzir suas proprias filhas, formando uma arvore de profundidade arbitraria.

Essa estrutura em arvore da as carteiras poder organizacional. Diferentes ramos da arvore servem a diferentes propositos, e a arvore inteira e recuperavel a partir da semente original.
