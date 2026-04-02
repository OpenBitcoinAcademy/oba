## BitVM: Verificacao de Computacao

BitVM habilita verificacao de computacao Turing-completa no Bitcoin sem mudancas de consenso. Um provador afirma que uma funcao f(x) = y. Se a afirmacao estiver correta, nada acontece on-chain. Se incorreta, qualquer verificador pode submeter uma prova de fraude e o provador perde uma garantia.

BitVM2 (2024-2025) tornou a verificacao sem permissao: qualquer pessoa pode contestar uma afirmacao falsa, e disputas se resolvem em tres transacoes on-chain. O sistema implementa um verificador SNARK em Bitcoin Script, habilitando verificacao eficiente de computacoes complexas.

Aplicacoes incluem bridges sem confianca (movendo bitcoin entre cadeias), rollups (agrupando transacoes off-chain com verificacao on-chain) e qualquer protocolo que se beneficie de "verificar, nao computar" no Bitcoin.
