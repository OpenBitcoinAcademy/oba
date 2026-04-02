## MuSig2 como Chave Interna

A chave interna P pode ser um agregado MuSig2 de multiplas chaves publicas. Se Alice e Bob agregam suas chaves em P usando MuSig2, a chave ajustada Q se compromete com ambos. Quando cooperam, produzem uma unica assinatura Schnorr em Q.

Na cadeia, esse multisig 2-de-2 parece identico a uma transacao de assinante unico. O output tem 34 bytes. A witness tem 64 bytes. Nenhum observador pode determinar que duas partes estavam envolvidas.

A BitGo relatou que um input de key path MuSig2 custa 57,5 bytes virtuais, comparado a 104,5 vbytes para um input multisig P2WSH SegWit nativo. A economia vem da eliminacao das chaves publicas e assinaturas por assinante que scripts multisig exigem na cadeia.
