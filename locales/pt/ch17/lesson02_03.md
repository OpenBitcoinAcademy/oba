## De Shares a uma Assinatura Padrao

O coordenador coleta t shares de assinatura e os soma. O resultado e uma unica assinatura Schnorr: um par (R, s) de 64 bytes, onde R e o ponto de nonce agregado e s e a soma dos shares de assinatura. Essa assinatura e valida sob a chave publica agregada do grupo.

Um no completo Bitcoin verificando essa transacao executa a equacao padrao de verificacao Schnorr BIP 340. Ele verifica uma assinatura contra uma chave publica. O no nao tem como saber se a assinatura foi produzida por tres assinantes, cinco assinantes ou um assinante. A verificacao e identica.

Por isso o FROST e poderoso para Bitcoin: nao requer mudancas de consenso. Taproot e BIP 340 ja aceitam as assinaturas que o FROST produz. A complexidade da assinatura threshold vive inteiramente no software dos assinantes. O blockchain e todos os verificadores permanecem inconscientes.
