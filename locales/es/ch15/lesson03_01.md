## Que son los descriptores de salida

Un descriptor de salida es una cadena que describe completamente como derivar un conjunto de direcciones de Bitcoin. Los BIPs 380 a 389 definen el lenguaje de descriptores. Cada descriptor especifica el tipo de script, las claves involucradas y las rutas de derivacion.

Un descriptor como `wpkh([d34db33f/84h/0h/0h]xpub.../0/*)` le dice a un monedero todo lo que necesita: usar P2WPKH, derivar desde esta clave publica extendida en esta ruta, y generar direcciones secuenciales. El monedero no necesita adivinar el tipo de script o el esquema de derivacion. El descriptor es autocontenido.

Antes de los descriptores, los monederos dependian de convenciones. BIP 44 decia que los monederos HD debian usar una ruta de derivacion especifica. BIP 49 agrego una ruta diferente para P2SH-SegWit. BIP 84 agrego otra para SegWit nativo. Un monedero que importaba un xpub tenia que probar todas las convenciones y esperar haber adivinado correctamente. Los descriptores reemplazan las conjeturas con declaraciones explicitas.
