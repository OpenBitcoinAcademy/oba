## Validacao do Lado do Cliente

No RGB, apenas as partes de um contrato validam suas transicoes de estado. Um emissor de tokens e os detentores dos tokens verificam o historico de transferencias. Nenhum outro no na rede precisa processar ou armazenar esses dados.

Compare com sistemas onde cada no valida cada contrato (o modelo do Ethereum). Validacao do lado do cliente escala sem limite: adicionar mais contratos nao aumenta a carga em nenhum no que nao esta envolvido. Privacidade e inerente: dados de contrato sao visiveis apenas para os participantes.

O custo: participantes devem armazenar e verificar o historico completo de seus ativos. Se o historico for perdido, o ativo nao pode ser verificado. RGB usa cadeias de commitment e provas para tornar essa verificacao eficiente.
