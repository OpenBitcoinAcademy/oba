## Regras de Validacao de Bloco

Quando um no recebe um novo bloco, ele verifica dezenas de criterios antes de aceita-lo. O hash do cabecalho do bloco deve ser menor que o alvo atual. O timestamp deve estar dentro de limites aceitaveis. A primeira transacao deve ser uma coinbase, e o output da coinbase nao deve exceder o subsidio permitido mais taxas. Cada outra transacao deve ser valida de acordo com as regras de script.

A raiz merkle do bloco deve corresponder a arvore de hash de todas as transacoes incluidas. O tamanho do bloco nao deve exceder o limite de consenso. O bloco deve referenciar um bloco pai valido que o no ja possui.

Se qualquer verificacao falhar, o no rejeita o bloco e pode desconectar do par que o enviou. Nao ha processo de recurso. Um bloco e valido ou nao e. Essa validacao rigorosa e o que impede mineradores de criar bitcoin do nada ou gastar moedas que nao possuem.
