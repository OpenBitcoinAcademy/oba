## Un nuevo lenguaje de scripts

Tapscript (BIP 342) define las reglas de validacion para scripts ejecutados dentro de gastos por ruta de script de Taproot. Es similar al Bitcoin Script heredado pero con mejoras especificas.

Todos los opcodes de verificacion de firma (OP_CHECKSIG, OP_CHECKSIGVERIFY) usan firmas Schnorr en lugar de ECDSA. Un nuevo opcode, OP_CHECKSIGADD, reemplaza al heredado OP_CHECKMULTISIG. En lugar de verificar multiples firmas contra multiples claves en una operacion, OP_CHECKSIGADD verifica una firma a la vez y acumula un contador. Esto habilita la verificacion por lotes: el verificador recopila todas las firmas y las verifica juntas en una sola operacion, mas rapido que verificar cada una individualmente.
