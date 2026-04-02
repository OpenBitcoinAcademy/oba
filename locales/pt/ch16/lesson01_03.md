## Por que Hardware Wallets Precisavam Disso

Hardware wallets tem uma restricao especifica: nao podem consultar o blockchain. Recebem dados, verificam o que podem, assinam e retornam o resultado. Sem um formato padrao, cada fabricante de hardware wallet tinha que escrever integracoes personalizadas para cada software de carteira. Adicionar uma nova hardware wallet significava corrigir cada coordenador. Adicionar um novo coordenador significava corrigir cada hardware wallet. A explosao combinatoria era insustentavel.

O PSBT resolveu isso definindo um formato que carrega toda informacao que um assinante precisa. O UTXO sendo gasto, o caminho de derivacao da chave, o tipo de sighash a usar, o redeem script para P2SH, o witness script para P2WSH. Uma hardware wallet recebe um PSBT, le os campos necessarios, assina, escreve sua assinatura parcial de volta no PSBT e o retorna. Sem protocolo proprietario. Sem negociacao de formato.

O ecossistema convergiu rapidamente. Coldcard, Trezor, Ledger, BitBox e Jade todos falam PSBT. Coordenadores de software como Sparrow, Specter e Bitcoin Core todos produzem e consomem PSBTs. Um quorum multisig pode usar hardware de diferentes fabricantes sem preocupacoes de compatibilidade.
