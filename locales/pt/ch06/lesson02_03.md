## Taxas de Transacao

Transacoes Bitcoin nao tem um campo explicito de taxa. A taxa e implicita: e a diferenca entre o valor total de todos os inputs e o valor total de todos os outputs.

Se seus inputs totalizam 100.000 satoshis e seus outputs totalizam 99.800 satoshis, a taxa e 200 satoshis. Mineradores coletam essa taxa como incentivo para incluir sua transacao em um bloco.

Taxas mais altas significam confirmacao mais rapida. Quando a rede esta ocupada, mineradores priorizam transacoes com taxas por unidade de peso mais altas. Uma transacao que paga muito pouco pode esperar horas ou dias.

A taxa depende do peso da transacao (medido em bytes virtuais, ou vbytes), nao do valor enviado. O SegWit introduziu unidades de peso onde dados de witness (assinaturas) recebem desconto comparado a outros dados da transacao. Uma transacao enviando 0,001 BTC pode custar a mesma taxa que uma enviando 1.000 BTC, se ambas tiverem a mesma estrutura.
