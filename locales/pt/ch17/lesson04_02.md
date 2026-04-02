## Abortos Identificaveis e MuSig2 vs FROST

Um participante malicioso pode interromper a assinatura enviando um share de assinatura invalido. Sem verificacoes adicionais, o coordenador somaria os shares e produziria uma assinatura final invalida, mas nao saberia qual assinante se comportou mal. FROST suporta abortos identificaveis: o coordenador pode verificar cada share de assinatura individualmente contra o share de chave publica do assinante. O assinante malicioso e identificado e pode ser excluido de sessoes futuras.

MuSig2 e FROST servem necessidades diferentes. MuSig2 e um esquema n-de-n: todos os participantes devem assinar, nao ha threshold, e o protocolo e mais simples. FROST e um esquema t-de-n: tolera assinantes ausentes, mas requer uma fase de geracao de chaves mais complexa. Ambos produzem output on-chain identico: uma unica assinatura Schnorr de 64 bytes sob uma chave publica de 32 bytes.

MuSig2 e adequado para cenarios onde todos os assinantes devem estar disponiveis, como um canal Lightning entre dois nos. FROST e adequado para arranjos de custodia onde redundancia importa, como um tesouro corporativo com cinco detentores de chave onde quaisquer tres podem autorizar um gasto.
