## O Algoritmo ECDSA

ECDSA (Elliptic Curve Digital Signature Algorithm) foi o esquema de assinatura original do Bitcoin. Ele opera na curva eliptica secp256k1, a mesma curva usada para geracao de chaves.

Para assinar uma mensagem, o algoritmo pega a chave privada e um hash da mensagem, combina-os com um numero aleatorio (chamado k, ou nonce), e produz dois valores: r e s. Juntos, (r, s) formam a assinatura.

Para verificar, o algoritmo pega a chave publica, o hash da mensagem e a assinatura (r, s). Ele realiza matematica de curva eliptica para verificar se a assinatura e consistente com a chave publica e a mensagem. Nenhuma chave privada e necessaria para verificacao.
