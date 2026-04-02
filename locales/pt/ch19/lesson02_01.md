## Selos de Uso Unico

RGB define um "selo" como um UTXO do Bitcoin. Quando uma transicao de estado ocorre (transferindo um token, atualizando um contrato), o selo e "fechado" gastando aquele UTXO. Um UTXO so pode ser gasto uma vez, entao um selo so pode ser fechado uma vez. Isso previne gasto duplo de ativos RGB usando o mecanismo de consenso existente do Bitcoin.

Os dados de transicao de estado (quem possui o que, atualizacoes de contrato) nunca tocam o blockchain. Apenas um commitment criptografico da transicao e incorporado na transacao Bitcoin, tipicamente em um output Taproot. O blockchain ancora o timing e a ordenacao. Os dados vivem off-chain com as partes envolvidas.
