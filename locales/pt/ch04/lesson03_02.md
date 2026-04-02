## As Etapas de Derivacao

Criar um endereco Bitcoin a partir de uma chave publica exige varias etapas de hashing.

Primeiro, faca o hash da chave publica com SHA-256. Depois, faca o hash desse resultado com RIPEMD-160. Isso produz um hash de 20 bytes chamado hash da chave publica, ou HASH160.

Em seguida, adicione um byte de versao na frente. Esse byte identifica a rede (mainnet vs testnet) e o tipo de endereco.

Por fim, codifique o resultado. Enderecos legacy usam codificacao Base58Check, que adiciona um checksum de 4 bytes e codifica em um alfabeto legivel que evita caracteres confusos como 0/O e l/1. Enderecos SegWit mais recentes usam codificacao Bech32, que e somente minusculas e tem deteccao de erros mais forte.
