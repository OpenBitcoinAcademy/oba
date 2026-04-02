## Tipos de Nos

Nem todo no executa o mesmo software ou desempenha a mesma funcao. Os nos diferem nos dados que armazenam e nas funcoes que fornecem.

Um **no completo** baixa e valida cada bloco e cada transacao contra as regras de consenso. Ele nao confia em ninguem. Pode verificar qualquer pagamento de forma independente. O Bitcoin Core e a implementacao de no completo mais comum.

Um **no de mineracao** e um no completo que tambem compete para criar novos blocos. Ele monta blocos candidatos a partir do mempool e executa proof of work. Nos de mineracao ganham o subsidio de bloco e as taxas de transacao quando encontram um bloco valido.

Um **no leve** (tambem chamado de cliente SPV) nao baixa blocos completos. Ele baixa apenas cabecalhos de bloco e solicita prova de que transacoes especificas existem. Nos leves confiam em nos completos ate certo ponto, trocando seguranca por menor uso de recursos.
