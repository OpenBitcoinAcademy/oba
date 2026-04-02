## Clientes SPV

Nem todo dispositivo pode armazenar e validar o blockchain completo. Celulares, dispositivos embarcados e hardware wallets tem armazenamento e poder de processamento limitados. Esses dispositivos executam clientes **Simplified Payment Verification** (SPV).

Um cliente SPV baixa apenas cabecalhos de bloco, nao blocos completos. A cadeia de cabecalhos e pequena: cerca de 60 MB para todo o historico do Bitcoin. O cliente pode verificar que um cabecalho de bloco atende ao alvo de proof of work, confirmando que um minerador gastou energia real para produzi-lo.

Para verificar se uma transacao esta confirmada, o cliente SPV solicita uma **prova Merkle** de um no completo. A prova mostra que o hash da transacao esta incluido na raiz Merkle de um bloco. O cliente confia que os nos completos validaram as transacoes do bloco, porque forjar um cabecalho com proof of work valido e proibitivamente caro.

SPV da aos clientes leves um nivel razoavel de seguranca sem baixar o blockchain completo. A troca e que um cliente SPV nao pode detectar se um bloco contem uma transacao invalida. Ele confia no peso economico do proof of work em vez de verificar todas as regras por conta propria.
