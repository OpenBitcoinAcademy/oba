## Politicas de gasto con decaimiento temporal

Una politica de recuperacion con timelock comienza con una condicion de gasto primaria y agrega condiciones de respaldo que se activan tras un retraso. La clave primaria controla los fondos de inmediato. Si la clave primaria se pierde o su titular queda indisponible, una ruta de gasto alternativa se desbloquea despues de un numero especificado de bloques.

En Miniscript, esto se expresa como: `or(pk(primary),and(pk(recovery),older(26280)))`. La clave primaria puede gastar en cualquier momento. La clave de recuperacion puede gastar solo despues de aproximadamente seis meses (26,280 bloques a 10 minutos por bloque). La politica se compila en un solo Script con dos rutas de ejecucion.

Este patron resuelve un problema real. Un monedero de una sola clave no tiene ruta de recuperacion. Si la clave se pierde, los fondos desaparecen. Un multisig estandar requiere multiples firmantes para cada transaccion, lo que agrega friccion al uso diario. Un respaldo con timelock le da al titular primario control total durante la operacion normal, y al mismo tiempo garantiza que una parte de confianza pueda recuperar los fondos tras un retraso.
