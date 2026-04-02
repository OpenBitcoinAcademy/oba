## Schnorr vs ECDSA

Ambos os algoritmos usam a curva secp256k1. Ambos fornecem seguranca equivalente contra recuperacao de chave. As diferencas sao praticas.

Assinaturas Schnorr sao menores (64 bytes vs 71-73). Schnorr suporta agregacao nativa de assinaturas. Schnorr tem uma equacao de verificacao mais simples. Schnorr tem uma prova formal de seguranca.

O ECDSA foi escolhido para o Bitcoin em 2008 porque o Schnorr era patenteado ate 2008, e o status da patente era incerto quando Satoshi projetou o sistema. Com a patente expirada e o Taproot ativado, novas transacoes podem usar Schnorr. O ECDSA continua suportado para compatibilidade retroativa.
