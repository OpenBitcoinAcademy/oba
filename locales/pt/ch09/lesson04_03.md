## Fee Sniping

Fee sniping e um ataque teorico onde um minerador, em vez de construir sobre o ultimo bloco, re-minera um bloco anterior para roubar suas taxas de transacao. Se as taxas no bloco N forem altas o suficiente, um minerador pode achar mais lucrativo fazer um fork da cadeia no bloco N-1 e coletar essas taxas para si.

O Bitcoin Core se defende contra isso definindo o locktime de novas transacoes para a altura de bloco atual. Uma transacao bloqueada para o bloco 800.000 nao pode ser incluida no bloco 799.999. Isso torna menos lucrativo re-minerar blocos antigos, porque as novas transacoes criadas desde entao estariam indisponiveis.

Fee sniping nao ocorreu na rede Bitcoin. A defesa e preventiva.
