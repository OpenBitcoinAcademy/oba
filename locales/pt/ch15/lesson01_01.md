## Escrevendo Bitcoin Script a Mao

Bitcoin Script e uma linguagem baseada em pilha. Cada condicao de gasto e uma sequencia de opcodes: empurrar uma chave, verificar uma assinatura, verificar um hash, aplicar um timelock. Para um gasto de chave unica, o script e curto e bem compreendido.

Combine condicoes, e a dificuldade cresce rapido. Um multisig 2-de-3 com alternativa de timelock requer ordenacao cuidadosa de ramos OP_IF, OP_CHECKMULTISIG e OP_CHECKSEQUENCEVERIFY. Um opcode mal posicionado muda completamente as condicoes de gasto. Uma chave duplicada pode passar despercebida ate que fundos fiquem bloqueados.

Nenhuma ferramenta pode analisar um Script arbitrario e responder: "Quem pode gastar isto, sob quais condicoes?" A linguagem carece de estrutura. Sequencias diferentes de opcodes podem codificar a mesma politica, e nao ha metodo geral para compara-las.
