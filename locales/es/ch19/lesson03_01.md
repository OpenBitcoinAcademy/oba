## Fedimint: custodia comunitaria

Fedimint es un protocolo para custodia operada por la comunidad usando ecash chaumiano federado. Un grupo de guardianes (miembros de confianza de la comunidad) operan conjuntamente una casa de moneda usando multisig umbral. Los usuarios depositan bitcoin y reciben billetes de ecash a cambio.

Los billetes de ecash se crean usando firmas ciegas. Al acunar, el usuario ciega un identificador aleatorio de token antes de enviarlo a los guardianes. Los guardianes lo firman sin ver el identificador. Cuando el usuario desciega el token firmado, la casa de moneda no puede vincular el deposito con el canje. Las transacciones entre usuarios no son rastreables.

Los usuarios confian en el umbral de la federacion (una mayoria de guardianes debe confabularse para robar). Esto es mas debil que el modelo sin confianza de Bitcoin pero mucho mas fuerte que confiar en un solo custodio. Fedimint apunta a comunidades con confianza social existente: grupos de ahorro, cooperativas, organizaciones locales.
