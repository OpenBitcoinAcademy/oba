## De n-de-n para t-de-n

MuSig2 tem uma limitacao: requer que todos os n participantes assinem. Cada detentor de chave deve estar online e cooperando no momento da assinatura. Se um participante perder sua chave ou ficar offline, os fundos ficam presos. Isso torna MuSig2 um esquema puramente n-de-n.

FROST (Flexible Round-Optimized Schnorr Threshold signatures) resolve isso. FROST e um esquema de assinatura threshold: quaisquer t de n participantes podem produzir uma assinatura valida. Uma configuracao FROST 3-de-5 significa que quaisquer tres dos cinco detentores de chave podem assinar. Os outros dois podem estar offline, indisponiveis ou ate permanentemente perdidos.

A assinatura que o FROST produz e uma assinatura Schnorr padrao de 64 bytes. On-chain, e indistinguivel de um gasto de chave unica. A politica de threshold, o numero de assinantes e as chaves publicas individuais estao todos ocultos. Um gasto FROST parece exatamente como um gasto por key path de um unico endereco Taproot.
