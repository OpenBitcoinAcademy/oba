## Fedimint: Custodia Comunitaria

Fedimint e um protocolo para custodia operada por comunidade usando ecash Chaumiano federado. Um grupo de guardioes (membros confiados da comunidade) opera conjuntamente um mint usando multisig threshold. Usuarios depositam bitcoin e recebem notas de ecash em troca.

As notas de ecash sao criadas usando assinaturas cegas. Ao cunhar, o usuario cega um identificador aleatorio de token antes de envia-lo aos guardioes. Os guardioes o assinam sem ver o identificador. Quando o usuario desvenda o token assinado, o mint nao pode vincular o deposito ao resgate. Transacoes entre usuarios sao desvinculaveis.

Usuarios confiam no threshold da federacao (uma maioria dos guardioes deve conspirar para roubar). Isso e mais fraco que o modelo sem confianca do Bitcoin, mas muito mais forte do que confiar em um unico custodiante. Fedimint mira comunidades com confianca social existente: grupos de poupanca, cooperativas, organizacoes locais.
