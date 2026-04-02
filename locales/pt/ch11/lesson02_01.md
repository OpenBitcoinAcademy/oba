## Estrutura do Cabecalho do Bloco

Todo bloco comeca com um cabecalho de 80 bytes contendo seis campos.

**Versao** (4 bytes): indica quais regras de consenso o bloco segue e sinaliza suporte a propostas de soft fork via versionbits.

**Hash do Bloco Anterior** (32 bytes): o hash double-SHA-256 do cabecalho do bloco anterior. Isso conecta cada bloco ao seu predecessor, formando a cadeia.

**Raiz Merkle** (32 bytes): um unico hash que se compromete com cada transacao no bloco. Mudar qualquer transacao muda a raiz merkle.

**Timestamp** (4 bytes): o horario aproximado em que o bloco foi minerado, em segundos Unix epoch.

**Target Bits** (4 bytes): uma codificacao compacta do alvo de proof of work. O hash do cabecalho do bloco deve ser menor que esse alvo.

**Nonce** (4 bytes): o valor que mineradores alteram para buscar um hash valido.
