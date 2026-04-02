## Geracao de Chaves Aleatorias

As primeiras carteiras Bitcoin geravam cada chave privada de forma independente usando um gerador de numeros aleatorios. Cada chave nao tinha relacao com as demais. A carteira armazenava todas em um arquivo de banco de dados.

Essa abordagem tinha um problema serio de backup. Cada nova chave exigia um novo backup. Se um usuario gerasse 100 chaves e fizesse backup apos a chave 50, as chaves 51 a 100 seriam perdidas se o arquivo da carteira fosse corrompido.

O Bitcoin Core originalmente pre-gerava um pool de 100 chaves para reduzir a frequencia de backups necessarios. Mas o problema permanecia: chaves aleatorias independentes sao dificeis de gerenciar.
