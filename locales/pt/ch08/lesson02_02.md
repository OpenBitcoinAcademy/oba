## Serializacao DER

Assinaturas ECDSA no Bitcoin sao codificadas usando DER (Distinguished Encoding Rules), um formato binario padrao. A codificacao DER envolve os valores r e s com bytes de tipo e comprimento.

Uma assinatura ECDSA tipica codificada em DER tem de 71 a 73 bytes. Ela e seguida por um flag SIGHASH de um byte que indica quais partes da transacao a assinatura cobre.

Transacoes segwit v0 ainda usam ECDSA, mas exigem codificacao DER estrita (BIP66). Isso eliminou uma fonte de maleabilidade de transacao onde a mesma assinatura valida podia ser codificada de multiplas formas, produzindo IDs de transacao diferentes.
