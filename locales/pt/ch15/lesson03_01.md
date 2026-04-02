## O que sao Output Descriptors

Um output descriptor e uma string que descreve completamente como derivar um conjunto de enderecos Bitcoin. Os BIPs 380 a 389 definem a linguagem de descriptors. Cada descriptor especifica o tipo de script, as chaves envolvidas e os caminhos de derivacao.

Um descriptor como `wpkh([d34db33f/84h/0h/0h]xpub.../0/*)` diz a uma carteira tudo que precisa: use P2WPKH, derive desta chave publica estendida neste caminho e gere enderecos sequenciais. A carteira nao precisa adivinhar o tipo de script ou esquema de derivacao. O descriptor e autocontido.

Antes dos descriptors, carteiras dependiam de convencoes. O BIP 44 dizia que carteiras HD deveriam usar um caminho de derivacao especifico. O BIP 49 adicionou um caminho diferente para P2SH-SegWit. O BIP 84 adicionou outro para SegWit nativo. Uma carteira importando um xpub tinha que tentar todas as convencoes e torcer para acertar. Descriptors substituem adivinhacao por declaracoes explicitas.
