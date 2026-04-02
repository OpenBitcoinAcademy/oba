## Assinaturas em Transacoes Bitcoin

No Bitcoin, assinaturas aparecem nos inputs de transacao (para transacoes legacy) ou na estrutura de witness (para transacoes segwit). Cada input que gasta um UTXO requer uma assinatura provando que o gastador controla a chave que bloqueou aquele output.

Multiplas partes podem colaborar em uma unica transacao. Cada parte assina apenas o input que controla. As assinaturas sao independentes: um assinante nao precisa de acesso a chave privada de outro.

O Bitcoin usa dois algoritmos de assinatura: ECDSA (Elliptic Curve Digital Signature Algorithm) para transacoes legacy e segwit v0, e assinaturas Schnorr para transacoes segwit v1 (Taproot).
