## Roteamento Cebola

O Lightning envolve instrucoes de pagamento em camadas de criptografia, como uma cebola. Cada no de encaminhamento descriptografa uma camada, que revela apenas o proximo salto e o valor a encaminhar. O no nao descobre o remetente, o destinatario final ou o comprimento total do caminho.

Esse modelo de privacidade (baseado em SPHINX) significa que Bob, encaminhando um pagamento de Alice para Carol, sabe apenas que Alice lhe enviou algo e que ele deve encaminhar para Carol. Ele nao sabe se Alice e a remetente original ou outro no de encaminhamento. Nao sabe se Carol e a destinataria final.

A cebola tem um tamanho fixo independente do numero de saltos, prevenindo analise de comprimento de caminho.
