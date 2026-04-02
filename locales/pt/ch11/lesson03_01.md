## O que e uma Arvore Merkle?

Uma arvore Merkle e uma arvore binaria de hashes. Cada no folha e o hash de uma transacao. Cada no interno e o hash da concatenacao de seus dois filhos. A raiz da arvore (a raiz Merkle) e um unico hash que se compromete com cada transacao no bloco.

Se um bloco contem quatro transacoes (A, B, C, D), a arvore e: hash(A) e hash(B) se combinam em hash(AB). hash(C) e hash(D) se combinam em hash(CD). hash(AB) e hash(CD) se combinam na raiz Merkle.

Se o numero de transacoes for impar, o ultimo hash e duplicado para fazer uma contagem par em cada nivel.
