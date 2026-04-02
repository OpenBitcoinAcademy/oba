## A Transacao Coinbase

Todo bloco comeca com uma transacao especial chamada transacao coinbase. Ela nao tem inputs de transacoes anteriores. Em vez disso, cria novos bitcoins como recompensa para o minerador que encontrou o bloco.

A transacao coinbase tem um input com uma referencia nula (nenhum output anterior sendo gasto). O minerador preenche o campo de script desse input com dados arbitrarios, incluindo o extra nonce usado durante a mineracao. Satoshi Nakamoto famosamente incorporou uma manchete de jornal na coinbase do bloco genesis: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."

Os outputs da transacao coinbase pagam o minerador. O valor total de output nao deve exceder o subsidio de bloco mais a soma de todas as taxas de transacao no bloco. Qualquer valor nao reivindicado e destruido.
