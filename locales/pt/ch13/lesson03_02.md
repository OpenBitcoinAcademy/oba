## Riscos de Consenso

Aplicacoes que validam transacoes devem implementar as regras de consenso de forma exata. Uma diferenca sutil em como duas implementacoes tratam um caso extremo pode fazer com que discordem sobre a validade de um bloco, dividindo a rede de sua perspectiva.

A abordagem mais segura para aplicacoes: delegar a validacao de consenso a um no completo (Bitcoin Core ou equivalente) e usar sua API para consultas ao blockchain. Nao reimplemente regras de consenso no codigo da aplicacao, a menos que esteja preparado para corresponder a cada caso extremo na implementacao de referencia.

Para desenvolvedores de carteiras: use bibliotecas bem testadas para derivacao de chaves, geracao de enderecos e assinatura de transacoes. Prefira implementacoes open-source testadas em batalha a codigo personalizado.
