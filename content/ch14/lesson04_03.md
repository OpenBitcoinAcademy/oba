## Future Upgrades via OP_SUCCESS

Tapscript reserves a set of opcodes called OP_SUCCESSx. If a script contains any OP_SUCCESS opcode, the entire script succeeds immediately without further validation.

This is a deliberate upgrade mechanism. A future soft fork can redefine an OP_SUCCESSx opcode to perform new validation (like OP_CHECKTEMPLATEVERIFY or OP_VAULT). Old nodes see OP_SUCCESS and accept the transaction. New nodes execute the redefined opcode and enforce the new rules.

Taproot also supports unknown leaf versions. The current Tapscript uses leaf version 0xC0. A future soft fork can define leaf version 0xC2 with entirely different script semantics. Old nodes skip validation of unknown leaf versions. New nodes enforce the new rules.

Both mechanisms allow Taproot to evolve without replacing it. Each script leaf version and each OP_SUCCESS opcode is a reserved slot for future functionality.
