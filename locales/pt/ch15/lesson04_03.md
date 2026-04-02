## Produtos Construidos com Miniscript

Liana Wallet e uma carteira Bitcoin open-source construida especificamente em torno da recuperacao com timelock do Miniscript. Desenvolvida pela Wizardsardine e lancada sob a licenca BSD 3-Clause, Liana permite que usuarios configurem uma chave de gasto primaria com um ou mais caminhos de recuperacao que ativam apos atrasos escolhidos. A carteira cuida da geracao de descriptor, derivacao de enderecos e construcao de PSBT. Usuarios definem sua politica por uma interface grafica sem escrever Miniscript a mao.

AnchorWatch adota uma abordagem diferente. Fundada por Rob Hamilton, a AnchorWatch fornece custodia de Bitcoin segurada pela Lloyd's of London. O modelo de custodia usa politicas Miniscript onde a AnchorWatch detem uma chave de recuperacao que ativa somente apos um timelock. O seguro cobre a perda da chave primaria. Como as condicoes de gasto sao codificadas em Miniscript, a seguradora pode verificar a politica on-chain: a chave de recuperacao nao tem acesso antes do timelock, e o detentor primario mantem controle total durante a operacao normal.

Ambos os produtos existem porque o Miniscript tornou politicas de gasto complexas portaveis e verificaveis. A politica esta no descriptor. Qualquer carteira, assinante ou auditor compativel pode parsea-la e confirmar as condicoes. A confianca esta na matematica, nao no fornecedor.
