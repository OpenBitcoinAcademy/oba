## Cashu: ecash con operador unico

Cashu es un protocolo de ecash chaumiano que usa un solo operador de casa de moneda en lugar de una federacion. Los usuarios depositan bitcoin (tipicamente via Lightning) y reciben tokens de ecash firmados de forma ciega. Los tokens son instrumentos al portador: quien posea un token valido puede canjearlo.

La casa de moneda no puede vincular la acunacion con el canje (las firmas ciegas lo aseguran). La casa de moneda no puede ver quien transfiere tokens a quien. Las transferencias entre usuarios son instantaneas y gratuitas (sin interaccion con la blockchain).

La contrapartida: los usuarios confian en un solo operador de casa de moneda. Si la casa de moneda es deshonesta o se desconecta, los fondos pueden perderse. Cashu esta disenado para despliegues a pequena escala y especificos de aplicacion (una cafeteria, un servicio de streaming, una comunidad) donde la relacion de confianza es manejable.
