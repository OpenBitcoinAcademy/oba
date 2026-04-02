## O Formato de Transmissao

Uma transacao Bitcoin e serializada em uma sequencia de bytes para transmissao pela rede e armazenamento em blocos. O formato de serializacao define a ordem exata de bytes e a codificacao de cada campo.

Uma transacao legacy tem quatro campos de nivel superior serializados em ordem: version (4 bytes, little-endian), inputs (variavel), outputs (variavel) e locktime (4 bytes, little-endian).

Uma transacao segwit adiciona tres campos: apos o version, um byte marker (0x00) e um byte flag (0x01) sinalizam que dados de witness seguem. A estrutura de witness aparece apos os outputs e antes do locktime. Um parser que ve um byte zero onde a contagem de inputs deveria estar sabe que esta lendo o formato estendido (segwit).
