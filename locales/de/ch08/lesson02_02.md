## DER-Serialisierung

ECDSA-Signatures in Bitcoin werden mit DER (Distinguished Encoding Rules) kodiert, einem standardisierten Binärformat. Die DER-Kodierung umschließt die r- und s-Werte mit Typ- und Längen-Bytes.

Eine typische DER-kodierte ECDSA-Signature ist 71 bis 73 Bytes lang. Ihr folgt ein Ein-Byte-SIGHASH-Flag, das angibt, welche Teile der Transaction die Signature abdeckt.

Segwit-v0-Transactions verwenden weiterhin ECDSA, verlangen aber strikte DER-Kodierung (BIP66). Das beseitigte eine Quelle von Transaction Malleability, bei der dieselbe gültige Signature auf mehrere Arten kodiert werden konnte und unterschiedliche Transaction-IDs erzeugte.
