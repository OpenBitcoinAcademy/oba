## SegWit Encapsulado (P2SH-P2WPKH)

Nem todas as carteiras suportavam enderecos segwit nativos (bc1...) quando o segwit foi ativado pela primeira vez. O segwit encapsulado fornece compatibilidade retroativa colocando o programa segwit dentro de um redeem script P2SH.

O output parece um endereco P2SH padrao (comecando com "3"). O redeem script contem: OP_0 <20-byte-pubkey-hash>. Ao gastar, o script de input revela o redeem script, e a witness fornece a assinatura e a chave publica.

O segwit encapsulado e um formato de transicao. O segwit nativo (enderecos bc1q...) e preferido para novas transacoes porque e menor, mais barato e tem melhor deteccao de erros via codificacao bech32.
