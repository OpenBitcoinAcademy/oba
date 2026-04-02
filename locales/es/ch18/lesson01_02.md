## Transacciones de compromiso

Cada parte posee su propia version de la transaccion de compromiso mas reciente. Estas transacciones dividen los fondos del canal segun el saldo actual. Si Alice tiene 0.7 BTC y Bob tiene 0.3 BTC, ambas transacciones de compromiso reflejan esta division.

Las transacciones de compromiso son asimetricas. La version de Alice paga a Bob de inmediato pero bloquea los fondos de Alice tras un timelock. La version de Bob hace lo opuesto. Esta asimetria habilita el mecanismo de penalizacion: si Alice transmite un compromiso antiguo (intentando reclamar mas de lo que le corresponde), Bob puede usar un secreto de revocacion para tomar todos los fondos antes de que expire el timelock de Alice.

Cada vez que el saldo se actualiza, ambas partes intercambian secretos de revocacion para los compromisos antiguos, haciendolos inseguros de transmitir.
