## Expandindo Multisig com o Tempo

Politicas com timelock podem expandir seu conjunto de assinantes ao longo do tempo. Uma empresa pode comecar com uma politica 2-de-2 entre dois cofundadores. Apos um ano, uma terceira chave (detida por um assessor juridico) ativa, e a politica se torna 2-de-3.

A expressao Miniscript e: `or(multi(2,founder_a,founder_b),and(multi(2,founder_a,founder_b,counsel),older(52560)))`. Durante o primeiro ano, apenas os dois fundadores podem assinar. Apos 52.560 blocos (aproximadamente um ano), quaisquer dois dos tres podem assinar.

A politica inteira vive em um unico UTXO. Nenhuma transacao on-chain e necessaria quando o timelock expira. O caminho de gasto adicional foi comprometido no momento da criacao do output. A chave do assessor ganha poder de gasto automaticamente quando a altura do bloco passa o limite. Isso remove a necessidade de uma cerimonia ativa de chaves no ponto de transicao.

Miniscript torna essas composicoes auditaveis. O compilador pode enumerar cada caminho de gasto e suas condicoes. Um revisor pode verificar que a chave do assessor nao tem poder de gasto antes do timelock, que os fundadores mantem controle total o tempo todo e que os tamanhos de witness ficam dentro dos limites de consenso.
