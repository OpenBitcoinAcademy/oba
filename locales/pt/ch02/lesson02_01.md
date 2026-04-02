## Alice Paga Bob

Alice quer comprar um produto na loja online de Bob. A pagina de checkout de Bob exibe uma opcao de pagamento em Bitcoin com um valor e um endereco Bitcoin.

Alice abre seu aplicativo de carteira. Ele varre o blockchain em busca de outputs de transacao nao gastos (UTXOs) vinculados aos enderecos dela. A carteira encontra um UTXO de 0,015 BTC, mais do que suficiente para cobrir o preco de Bob de 0,01 BTC.

A carteira de Alice constroi uma transacao com um input (referenciando seu UTXO de 0,015 BTC) e dois outputs: 0,01 BTC para o endereco de Bob e 0,0049 BTC de volta para Alice como troco. Os 0,0001 BTC restantes sao a taxa de transacao.
