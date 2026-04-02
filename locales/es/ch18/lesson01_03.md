## El mecanismo de revocacion

Cuando Alice y Bob actualizan el saldo de su canal, invalidan el compromiso anterior intercambiando secretos de revocacion. Cada parte ahora posee un secreto que puede castigar a la otra por transmitir el estado antiguo.

Si Bob transmite un compromiso antiguo donde tenia 0.5 BTC (pero el saldo actual le da solo 0.3 BTC), Alice puede usar el secreto de revocacion de Bob para reclamar el saldo completo del canal. Bob pierde todo.

Esta penalizacion hace que el fraude sea economicamente irracional. El costo de intentar un fraude (perder todos los fondos) supera con creces la ganancia potencial (reclamar un saldo antiguo ligeramente mayor). El comportamiento honesto es la estrategia dominante.
