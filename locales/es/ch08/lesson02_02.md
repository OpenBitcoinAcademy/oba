## Serialización DER

Las firmas ECDSA en Bitcoin se codifican usando DER (Distinguished Encoding Rules), un formato binario estándar. La codificación DER envuelve los valores r y s con bytes de tipo y longitud.

Una firma ECDSA codificada en DER típica tiene de 71 a 73 bytes de longitud. Va seguida de un byte de SIGHASH que indica qué partes de la transacción cubre la firma.

Las transacciones segwit v0 todavía usan ECDSA pero requieren codificación DER estricta (BIP66). Esto eliminó una fuente de maleabilidad de transacciones donde la misma firma válida podía codificarse de múltiples formas, produciendo diferentes IDs de transacción.
