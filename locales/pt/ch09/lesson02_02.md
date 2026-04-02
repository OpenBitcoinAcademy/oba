## Unidades de Peso

O SegWit introduziu unidades de peso para substituir a metrica de tamanho baseada em bytes. Cada byte de dados nao-witness custa 4 unidades de peso. Cada byte de dados de witness (assinaturas) custa 1 unidade de peso.

Esse desconto incentiva o uso de transacoes segwit, que armazenam assinaturas na estrutura de witness. Uma transacao segwit usa menos peso que uma transacao legacy com o mesmo numero de inputs e outputs.

Bytes virtuais (vbytes) convertem peso para equivalente em bytes: vbytes = peso / 4. Uma transacao com 600 unidades de peso tem 150 vbytes. Taxas expressas em sat/vB contabilizam o desconto segwit automaticamente.
