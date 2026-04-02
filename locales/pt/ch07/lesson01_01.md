## Scripts de Bloqueio

Cada output de transacao contem um script de bloqueio, chamado scriptPubKey. Esse script define as condicoes que devem ser atendidas para gastar o output.

O script de bloqueio mais simples de entender e o Pay-to-Public-Key-Hash (P2PKH). Ele diz: "Para gastar este output, prove que voce possui a chave privada correspondente a este hash de chave publica." Transacoes modernas usam formatos mais novos (P2WPKH para SegWit, P2TR para Taproot), mas o P2PKH demonstra o conceito central com mais clareza.

Em notacao de script, um bloqueio P2PKH se parece com: OP_DUP OP_HASH160 <pubkey_hash> OP_EQUALVERIFY OP_CHECKSIG. Cada parte e uma instrucao que o motor de script do Bitcoin executa.
