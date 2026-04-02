## Duas Transacoes para Milhares de Pagamentos

Um canal de pagamento e um contrato multisig 2-de-2 no blockchain do Bitcoin. Duas partes bloqueiam fundos em um output compartilhado. Elas trocam transacoes de commitment assinadas off-chain para atualizar a divisao do saldo. Apenas duas transacoes on-chain sao necessarias: uma para abrir o canal e uma para fechar.

Entre abertura e fechamento, as partes podem trocar milhares de pagamentos. Cada pagamento atualiza a transacao de commitment que divide os fundos bloqueados. O blockchain nunca ve esses estados intermediarios. O resultado: pagamentos rapidos (milissegundos), taxas quase zero e sem atraso de confirmacao de bloco.
