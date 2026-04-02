## SegWit Nativo (P2WPKH e P2WSH)

O Segregated Witness moveu dados de assinatura do script de input para uma estrutura de witness separada. Isso corrigiu a maleabilidade de transacao (terceiros podiam modificar o txid alterando a codificacao da assinatura) e habilitou o desconto de witness para taxas.

P2WPKH (Pay to Witness Public Key Hash) e o equivalente segwit do P2PKH. O script de output contem: OP_0 <20-byte-pubkey-hash>. A witness fornece a assinatura e a chave publica. O script de input esta vazio.

P2WSH (Pay to Witness Script Hash) e o equivalente segwit do P2SH. O script de output contem: OP_0 <32-byte-script-hash>. A witness fornece o script e seus dados necessarios. O P2WSH usa um hash SHA-256 de 32 bytes em vez do HASH160 de 20 bytes do P2SH, fornecendo resistencia a colisao mais forte.
