## O Problema Antes dos PSBTs

Uma transacao Bitcoin precisa de informacoes de multiplas fontes antes de poder ser transmitida. A carteira que cria a transacao deve saber quais UTXOs gastar. O assinante deve possuir as chaves privadas. Em muitas configuracoes, esses sao dispositivos ou softwares diferentes.

Antes do BIP 174, nao havia forma padrao de passar uma transacao nao assinada entre esses participantes. Cada software de carteira inventava seu proprio formato. O Bitcoin Core serializava transacoes parciais de forma diferente do Electrum, que serializava de forma diferente das hardware wallets. Uma transacao criada em uma ferramenta nao podia ser assinada por outra sem codigo de integracao personalizado.

Isso tornava configuracoes multi-assinatura dolorosas. Cada assinante precisava de software compativel. Coordenar entre um laptop, uma hardware wallet e uma maquina de cold storage exigia passos manuais e conversoes de formato que introduziam erros.
