## Transacoes de Commitment

Cada parte detem sua propria versao da ultima transacao de commitment. Essas transacoes dividem os fundos do canal de acordo com o saldo atual. Se Alice tem 0,7 BTC e Bob tem 0,3 BTC, ambas as transacoes de commitment refletem essa divisao.

Transacoes de commitment sao assimetricas. A versao de Alice paga Bob imediatamente, mas bloqueia os fundos de Alice atras de um timelock. A versao de Bob faz o oposto. Essa assimetria habilita o mecanismo de penalidade: se Alice transmitir um commitment antigo (tentando reivindicar mais do que possui), Bob pode usar um segredo de revogacao para tomar todos os fundos antes que o timelock de Alice expire.

Cada vez que o saldo atualiza, ambas as partes trocam segredos de revogacao dos commitments antigos, tornando-os inseguros para transmitir.
