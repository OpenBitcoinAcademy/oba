## Uma Nova Linguagem de Script

Tapscript (BIP 342) define as regras de validacao para scripts executados dentro de gastos por script path do Taproot. E similar ao Bitcoin Script legacy, mas com melhorias direcionadas.

Todos os opcodes de verificacao de assinatura (OP_CHECKSIG, OP_CHECKSIGVERIFY) usam assinaturas Schnorr em vez de ECDSA. Um novo opcode, OP_CHECKSIGADD, substitui o legado OP_CHECKMULTISIG. Em vez de verificar multiplas assinaturas contra multiplas chaves em uma operacao, OP_CHECKSIGADD verifica uma assinatura por vez e acumula um contador. Isso habilita verificacao em lote: o verificador coleta todas as assinaturas e as verifica juntas em uma unica operacao, mais rapido do que verificar cada uma individualmente.
