## Updater e Signer

O Updater adiciona os metadados que Signers precisam. Para cada input, o Updater anexa o UTXO sendo gasto (a transacao anterior completa ou o witness UTXO especifico), o caminho de derivacao BIP 32 da chave de assinatura, o redeem script se o input for P2SH e o witness script se o input for P2WSH. Para cada output, o Updater pode anexar caminhos de derivacao BIP 32 para que o assinante possa verificar que outputs de troco pertencem a mesma carteira.

O Signer le o PSBT, identifica os inputs que pode assinar e produz assinaturas parciais. Para cada input que assina, escreve uma entrada PARTIAL_SIG indexada pela chave publica. O Signer nao modifica a transacao em si. Nao finaliza scripts. Adiciona sua assinatura e passa o PSBT para o proximo participante.

Um Signer deve verificar os dados de UTXO antes de assinar. Se o PSBT alegar que um input gasta 0,5 BTC, mas o UTXO real contem 1,0 BTC, o assinante estaria sem saber doando 0,5 BTC em taxas. Hardware wallets verificam valores de UTXO contra valores de output e alertam o usuario se a taxa parecer excessiva.
