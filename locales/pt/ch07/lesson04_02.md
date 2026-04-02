## MAST e Tapscript

A arvore Merkle no Taproot e uma Merklized Alternative Script Tree (MAST). Cada folha da arvore e um script diferente (uma condicao de gasto diferente). Para gastar via script path, o gastador revela apenas o script que esta usando e fornece uma prova Merkle de que ele faz parte da arvore comprometida.

Scripts nao usados permanecem ocultos. Um output Taproot com 100 condicoes de gasto possiveis parece igual no blockchain a um com 1 condicao, a menos que o script path seja exercido. Ramos nao usados permanecem privados.

Tapscript (BIP342) e a linguagem de script usada dentro das folhas de script path do Taproot. E similar ao Script legacy, mas com melhorias: OP_CHECKSIGADD substitui OP_CHECKMULTISIG (permitindo validacao em lote), opcodes de assinatura usam Schnorr em vez de ECDSA, e o limite de tamanho de script e removido (substituido por limites de peso).
