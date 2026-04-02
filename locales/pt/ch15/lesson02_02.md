## O Sistema de Tipos

Cada fragmento Miniscript tem um tipo que descreve seu comportamento na pilha do Script. Os quatro tipos base sao B, V, K e W.

Uma expressao tipo B (base) empurra um unico valor diferente de zero em caso de sucesso e zero em caso de falha. Uma expressao tipo V (verify) ou tem sucesso sem saida na pilha ou aborta o script. Uma expressao tipo K (key) empurra uma chave publica em caso de sucesso. Uma expressao tipo W (wrapped) se comporta como B mas consome o elemento do topo da pilha primeiro, usada para combinar sub-expressoes.

O compilador verifica tipos em cada ponto de composicao. Um fragmento `and_v` requer que seu primeiro argumento seja tipo V e seu segundo seja tipo B. Se voce passar argumentos com tipos errados, o compilador rejeita a expressao. Isso captura erros que, em Script bruto, produziriam um script que "funciona" mas aplica condicoes erradas.

O sistema de tipos tambem rastreia propriedades como dissipabilidade (um ramo insatisfeito pode deixar a pilha limpa?) e nao-maleabilidade (existe exatamente uma witness valida para cada caminho de gasto?). Essas propriedades se propagam pela composicao. Uma expressao Miniscript as tem ou nao tem, e o compilador pode informar.
