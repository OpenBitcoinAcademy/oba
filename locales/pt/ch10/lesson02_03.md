## Sincronizando o Blockchain

Um novo no deve baixar e validar o blockchain inteiro antes de poder verificar transacoes atuais. Esse processo e chamado **Initial Block Download** (IBD).

O no envia mensagens `getheaders` para seus pares, solicitando cabecalhos de bloco em lotes. Cabecalhos sao pequenos (80 bytes cada) e chegam rapidamente. O no os usa para construir a cadeia de cabecalhos com mais proof of work cumulativo.

Uma vez que o no identifica a melhor cadeia de cabecalhos, ele solicita blocos completos usando mensagens `getdata`. Ele baixa blocos de multiplos pares em paralelo para acelerar o processo. Cada bloco e validado contra as regras de consenso conforme chega: assinaturas de transacao, execucao de script, limites de peso e o alvo de proof of work.

O IBD pode levar horas ou dias em hardware lento. Apos completar, o no muda para operacao normal. Ele recebe novos blocos e transacoes conforme sao transmitidos e os valida em tempo real.
