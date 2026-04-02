## Replace By Fee (RBF)

Se uma transacao esta presa no mempool porque sua taxa e muito baixa, o remetente pode substitui-la por uma nova versao que paga uma taxa mais alta. Isso e Replace By Fee (RBF), definido no BIP125.

Para usar RBF, a transacao original deve sinalizar substituibilidade definindo o numero de sequencia de pelo menos um input para um valor abaixo de 0xFFFFFFFE. O software de carteira cuida disso automaticamente.

A transacao substituta deve pagar uma taxa total mais alta que a original. Ela pode mudar os outputs (pagando um valor diferente ou adicionando/removendo destinatarios) desde que gaste pelo menos um dos mesmos inputs.
