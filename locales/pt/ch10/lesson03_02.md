## Compact Block Relay

Definido no **BIP 152**, o compact block relay reduz a largura de banda necessaria para propagar um novo bloco. O insight principal: a maioria das transacoes em um novo bloco ja esta no mempool do no receptor.

Em vez de enviar o bloco completo, o no anunciante envia uma mensagem `cmpctblock`. Ela contem o cabecalho do bloco, uma lista de IDs curtos de transacao e a transacao coinbase. IDs curtos de transacao sao hashes truncados de 6 bytes que permitem ao receptor combinar transacoes que ele ja possui.

O no receptor reconstroi o bloco a partir de seu proprio mempool usando os IDs curtos. Se alguma transacao estiver faltando, ele a solicita com uma mensagem `getblocktxn` e a recebe em uma resposta `blocktxn`.

O BIP 152 define dois modos. No **modo de baixa largura de banda**, um no primeiro anuncia o bloco com uma mensagem `inv`. O par solicita o bloco compacto apenas se interessado. No **modo de alta largura de banda**, o no envia o bloco compacto imediatamente sem esperar uma solicitacao. Mineradores e nos bem conectados tipicamente usam o modo de alta largura de banda para minimizar a latencia.

O compact block relay reduz a largura de banda de propagacao em 90% ou mais para um bloco tipico. Propagacao mais rapida da aos mineradores menores uma chance mais justa, fortalecendo a descentralizacao.
