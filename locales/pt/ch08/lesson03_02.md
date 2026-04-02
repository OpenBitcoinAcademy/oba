## Agregacao de Multi-assinatura

Com ECDSA, uma transacao que requer tres assinantes precisa de tres assinaturas separadas, cada uma de 71-73 bytes. As assinaturas sao verificadas individualmente.

Com Schnorr, os mesmos tres assinantes podem combinar suas assinaturas em uma unica assinatura agregada de 64 bytes usando protocolos como MuSig2. No blockchain, a assinatura agregada parece identica a uma assinatura de um unico assinante. Nenhum observador pode dizer quantas partes participaram.

Isso tem dois beneficios. Transacoes com multiplos assinantes usam menos espaco de bloco (taxas menores). E transacoes multi-assinatura se tornam indistinguiveis de transacoes de assinatura unica, melhorando a privacidade.
