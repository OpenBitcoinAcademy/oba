## Construir sistemas Bitcoin de forma segura

Los desarrolladores que construyen sobre Bitcoin enfrentan un desafio unico: los errores pueden perder dinero. Una falla en la generacion de claves, la construccion de transacciones o la validacion de firmas puede resultar en fondos robados o perdidos permanentemente.

El modelo de consenso descentralizado significa que no hay autoridad para revertir una transaccion erronea. El codigo que maneja claves privadas debe tratarlas como los datos mas sensibles del sistema. Las claves deben generarse a partir de fuentes de entropia de alta calidad, almacenarse cifradas en reposo y borrarse de la memoria despues de su uso.

Cada componente que toca claves o construye transacciones debe ser auditado, probado contra vectores conocidos y sometido a revision adversarial.
