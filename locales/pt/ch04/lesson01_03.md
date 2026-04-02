## SHA-256 no Bitcoin

O Bitcoin usa uma funcao hash chamada SHA-256, projetada pela Agencia de Seguranca Nacional dos Estados Unidos. SHA significa Secure Hash Algorithm. 256 refere-se ao tamanho da saida: 256 bits, ou 32 bytes.

O SHA-256 aparece em todo lugar no Bitcoin. IDs de transacao sao hashes double-SHA-256. A mineracao envolve encontrar entradas que produzam hashes abaixo de um valor-alvo. A derivacao de enderecos usa SHA-256 combinado com RIPEMD-160 (uma funcao hash diferente que produz uma saida mais curta de 20 bytes). A combinacao de ambos, chamada HASH160, e o que cria o hash da chave publica em um endereco.
