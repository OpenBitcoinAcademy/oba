## Campos por Input

Cada mapa de input descreve um input da transacao. Os campos criticos que um assinante precisa sao os dados de UTXO, o caminho de derivacao da chave de assinatura e quaisquer scripts necessarios para construir a witness.

WITNESS_UTXO (tipo de chave 0x01) armazena o output sendo gasto: o valor em satoshis e o scriptPubKey. Para inputs SegWit, isso e suficiente para assinar porque o algoritmo de sighash se compromete com o valor diretamente. NON_WITNESS_UTXO (tipo de chave 0x00) armazena a transacao anterior inteira, necessaria para inputs legacy onde o assinante deve verificar o valor consultando o output na transacao completa.

BIP32_DERIVATION (tipo de chave 0x06) mapeia uma chave publica para seu caminho de derivacao BIP 32 e fingerprint da chave-mestra. O assinante combina o fingerprint com sua propria chave-mestra, deriva a chave privada no caminho dado e assina. PARTIAL_SIG (tipo de chave 0x02) armazena uma assinatura indexada pela chave publica que a produziu. SIGHASH_TYPE (tipo de chave 0x03) especifica qual flag de sighash o assinante deve usar.

Para inputs P2SH, REDEEM_SCRIPT (tipo de chave 0x04) carrega o redeem script. Para inputs P2WSH, WITNESS_SCRIPT (tipo de chave 0x05) carrega o witness script. O assinante precisa desses para calcular o sighash correto.
