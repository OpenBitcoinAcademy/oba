## Tipos de SIGHASH

Uma assinatura nao se compromete com a transacao inteira por padrao. O flag SIGHASH anexado a cada assinatura especifica quais partes da transacao a assinatura cobre.

**SIGHASH_ALL** (0x01) e o padrao. A assinatura cobre todos os inputs e todos os outputs. Mudar qualquer parte da transacao invalida a assinatura.

**SIGHASH_NONE** (0x02) cobre todos os inputs, mas nenhum output. Qualquer pessoa pode mudar os outputs apos a assinatura. Usado em protocolos colaborativos raros.

**SIGHASH_SINGLE** (0x03) cobre todos os inputs, mas apenas o output no mesmo indice do input sendo assinado. Outros outputs podem ser alterados.

O modificador ANYONECANPAY pode ser combinado com qualquer um desses, permitindo assinaturas que cobrem apenas um unico input.
