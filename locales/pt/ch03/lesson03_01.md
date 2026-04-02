## A Interface JSON-RPC

O Bitcoin Core expoe uma interface JSON-RPC que permite que programas consultem dados do blockchain e enviem transacoes. A ferramenta de linha de comando `bitcoin-cli` envia requisicoes para essa interface.

Cada dado no blockchain e consultavel: cabecalhos de blocos, blocos completos, transacoes individuais, saldos de enderecos, conteudo do mempool, conexoes de pares e estatisticas de rede.

Carteiras, exploradores de blocos, processadores de pagamento e ferramentas de pesquisa se comunicam com o Bitcoin Core por essa API. A interface e a ponte entre os dados brutos do blockchain e os aplicativos que os apresentam aos usuarios.
