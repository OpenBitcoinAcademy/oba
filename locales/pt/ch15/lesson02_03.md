## Mapeamento Biunivoco

A relacao entre Miniscript e Script e biunivoca para o conjunto de fragmentos suportados. Cada expressao Miniscript mapeia para exatamente uma codificacao em Script, e cada padrao de Script suportado mapeia de volta para exatamente uma expressao Miniscript.

Considere o fragmento `pk(K)`: codifica para `<K> OP_CHECKSIG`. O fragmento `and_v(v:pk(A),pk(B))` codifica para `<A> OP_CHECKSIGVERIFY <B> OP_CHECKSIG`. Nao ha ambiguidade. Dado o Script, voce pode recuperar o Miniscript original.

Essa propriedade torna possivel a analise de ida e volta. Uma carteira pode receber um Script de outra parte, decodifica-lo para Miniscript e determinar as condicoes de gasto sem confiar na descricao do remetente. Duas carteiras construidas por equipes diferentes podem concordar em uma politica, compilar para Miniscript de forma independente e verificar que produziram o mesmo Script.

Nem todos os Bitcoin Scripts validos podem ser representados em Miniscript. O conjunto de fragmentos cobre os padroes necessarios para politicas de gasto praticas: chaves, hashes, timelocks, thresholds e combinacoes booleanas. Scripts que usam sequencias de opcodes incomuns ou manipulacao de pilha ficam fora do subconjunto Miniscript.
