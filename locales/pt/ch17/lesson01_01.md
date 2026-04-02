## O Panorama do Multisig

O Bitcoin suporta custodia multi-parte desde 2012 por meio do OP_CHECKMULTISIG. Um script multisig 2-de-3 lista todas as tres chaves publicas on-chain e exige duas assinaturas validas. Isso funciona, mas revela exatamente quantas partes controlam os fundos. Qualquer pessoa inspecionando o blockchain ve o threshold, o numero total de assinantes e cada chave publica envolvida.

O custo on-chain escala linearmente. Um multisig 3-de-5 coloca cinco chaves publicas de 33 bytes e tres assinaturas de 72 bytes na witness. Sao 381 bytes de dados de witness antes de contar o proprio script. Uma configuracao 7-de-10 custa ainda mais. Taxas sobem, privacidade cai, e a politica de gasto e visivel para todos os observadores.

Taproot e assinaturas Schnorr mudaram o que e possivel. MuSig2 permite que n participantes produzam uma unica assinatura agregada que parece identica a uma assinatura Schnorr individual on-chain. A chave publica combinada tem 32 bytes. A assinatura tem 64 bytes. Nenhum observador pode dizer quantas partes estavam envolvidas.
