## Actualizaciones futuras mediante OP_SUCCESS

Tapscript reserva un conjunto de opcodes llamados OP_SUCCESSx. Si un script contiene cualquier opcode OP_SUCCESS, el script completo tiene exito inmediatamente sin mas validacion.

Este es un mecanismo de actualizacion deliberado. Un futuro soft fork puede redefinir un opcode OP_SUCCESSx para realizar nueva validacion (como OP_CHECKTEMPLATEVERIFY u OP_VAULT). Los nodos antiguos ven OP_SUCCESS y aceptan la transaccion. Los nodos nuevos ejecutan el opcode redefinido y aplican las nuevas reglas.

Taproot tambien soporta versiones de hoja desconocidas. El Tapscript actual usa la version de hoja 0xC0. Un futuro soft fork puede definir la version de hoja 0xC2 con semantica de script completamente diferente. Los nodos antiguos omiten la validacion de versiones de hoja desconocidas. Los nodos nuevos aplican las nuevas reglas.

Ambos mecanismos permiten a Taproot evolucionar sin reemplazarlo. Cada version de hoja de script y cada opcode OP_SUCCESS es un espacio reservado para funcionalidad futura.
