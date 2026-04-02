## Gastando com a Chave Ajustada

O gasto por key path e a forma mais simples e mais privada de gastar um output Taproot. A witness contem um unico elemento: uma assinatura Schnorr BIP 340 de 64 bytes para a chave ajustada Q.

O gastador calcula a chave privada ajustada (chave privada original + tweak), assina a transacao e fornece a assinatura. Nenhum script e revelado. Nenhuma prova Merkle e necessaria. A witness inteira tem 64 bytes.

Um verificador checa a assinatura contra Q (a chave no scriptPubKey do output). Se valida, o gasto esta autorizado. O verificador nao sabe nem se importa se Q foi derivado de uma unica chave, um agregado MuSig2 ou uma chave com arvore de scripts incorporada.
