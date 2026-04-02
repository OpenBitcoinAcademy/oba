## El tweak de seguridad de Taproot

ChillDKG produce una clave publica de grupo, pero esta clave no puede usarse directamente como clave de salida Taproot. BIP 341 requiere que la clave de salida Q sea una version con tweak de una clave interna P: Q = P + tagged_hash("TapTweak", P) * G para una salida de solo ruta de clave.

BIP 445 especifica que ChillDKG aplica este tweak como parte de la generacion de claves. El protocolo calcula la clave publica de grupo con tweak y ajusta las porciones de clave para que el grupo pueda firmar con la clave con tweak sin pasos adicionales durante la firma. La porcion de cada participante se modifica para tener en cuenta el tweak, y la clave publica de grupo que devuelve ChillDKG es la clave de salida Taproot final.

Esta integracion importa porque el tweak debe aplicarse de forma consistente. Si los participantes no estan de acuerdo en la clave con tweak, la firma fallara. Al integrar el tweak en el propio protocolo DKG, ChillDKG garantiza que todos los participantes deriven la misma clave de salida y posean porciones consistentes con ella.
