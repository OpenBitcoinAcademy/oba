## O Panorama Geral

Assinaturas threshold representam uma mudanca em como a custodia de Bitcoin funciona. A abordagem legada revelava politicas de gasto on-chain: qualquer pessoa podia contar as chaves, ver o threshold e rastrear carteiras multisig entre transacoes. FROST e ChillDKG movem toda essa complexidade para off-chain.

O grupo negocia seu threshold e gera chaves usando ChillDKG. Quaisquer t assinantes produzem uma assinatura usando o protocolo FROST. O blockchain ve um gasto por key path Taproot. Verificadores checam uma assinatura contra uma chave. O custo on-chain e constante independentemente de quantos participantes estao envolvidos: 64 bytes para a assinatura, 32 bytes para a chave.

Isso e possivel porque Taproot e verificacao Schnorr BIP 340 ja estao implantados na rede Bitcoin. Nenhum soft fork e necessario. Nenhum novo opcode. O protocolo de assinatura threshold roda inteiramente nas carteiras dos participantes. A camada de consenso verifica o resultado usando as mesmas regras que usa para cada outro gasto Taproot.
