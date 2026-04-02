## Atualizacoes Futuras via OP_SUCCESS

Tapscript reserva um conjunto de opcodes chamados OP_SUCCESSx. Se um script contem qualquer opcode OP_SUCCESS, o script inteiro tem sucesso imediatamente sem validacao adicional.

Esse e um mecanismo deliberado de atualizacao. Um futuro soft fork pode redefinir um opcode OP_SUCCESSx para executar nova validacao (como OP_CHECKTEMPLATEVERIFY ou OP_VAULT). Nos antigos veem OP_SUCCESS e aceitam a transacao. Nos novos executam o opcode redefinido e aplicam as novas regras.

Taproot tambem suporta versoes de folha desconhecidas. O Tapscript atual usa versao de folha 0xC0. Um futuro soft fork pode definir versao de folha 0xC2 com semanticas de script totalmente diferentes. Nos antigos pulam a validacao de versoes de folha desconhecidas. Nos novos aplicam as novas regras.

Ambos os mecanismos permitem que o Taproot evolua sem substitui-lo. Cada versao de folha de script e cada opcode OP_SUCCESS e um slot reservado para funcionalidade futura.
