## O Requisito de Aleatoriedade

Tanto ECDSA quanto Schnorr exigem um numero aleatorio (nonce) durante a criacao da assinatura. A seguranca de todo o sistema depende dessa aleatoriedade.

Se o mesmo nonce for usado para assinar duas mensagens diferentes com a mesma chave privada, a chave privada pode ser calculada a partir das duas assinaturas. Isso nao e um risco teorico. Em 2013, uma falha no gerador de numeros aleatorios do Android fez multiplas carteiras Bitcoin reutilizarem nonces, e fundos foram roubados.

Implementacoes modernas usam nonces deterministicos (RFC 6979 para ECDSA, BIP340 para Schnorr), derivando o nonce da chave privada e do hash da mensagem. Isso elimina a dependencia de um gerador de numeros aleatorios e previne reutilizacao de nonce.
