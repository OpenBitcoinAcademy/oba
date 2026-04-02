## Combiner, Finalizer, Extractor

O Combiner mescla multiplos PSBTs que contem assinaturas parciais para a mesma transacao. Em um multisig 2-de-3, o assinante A produz um PSBT com sua assinatura e o assinante B produz outro. O Combiner pega ambos, mescla as entradas PARTIAL_SIG de cada input e produz um unico PSBT com todas as assinaturas disponiveis.

O Finalizer transforma assinaturas parciais em scriptSig e witness completos para cada input. Para um input P2PKH, ele pega a unica PARTIAL_SIG e constroi o scriptSig. Para um multisig P2WSH 2-de-3, ele pega as assinaturas parciais, as ordena corretamente e constroi a pilha de witness com o redeem script. Apos a finalizacao, os mapas de input do PSBT contem campos FINAL_SCRIPTSIG e FINAL_SCRIPTWITNESS. Os dados parciais nao sao mais necessarios.

O Extractor le o PSBT finalizado e produz a transacao bruta de rede. Ele pega a transacao nao assinada do mapa global, insere o FINAL_SCRIPTSIG e FINAL_SCRIPTWITNESS de cada input e serializa o resultado. A saida e uma transacao Bitcoin padrao pronta para transmissao. O PSBT cumpriu seu proposito.
