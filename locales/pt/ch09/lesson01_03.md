## Transacoes Conflitantes

Alice pode criar duas transacoes que gastam o mesmo output: uma pagando Bob e outra pagando Carol. Ambas sao assinaturas validas, mas apenas uma pode ser incluida no blockchain (o output so pode ser gasto uma vez).

Mineradores decidem qual transacao conflitante incluir. Por padrao, a maioria dos mineradores segue uma politica de "primeiro visto" e inclui a transacao que recebeu primeiro. Mas isso e uma politica, nao uma regra de consenso. Um minerador pode incluir qualquer transacao valida.

Por isso Bob deveria esperar por confirmacoes antes de considerar um pagamento como final. A taxa que Alice paga incentiva mineradores a incluir sua transacao prontamente, reduzindo a janela para transacoes conflitantes.
