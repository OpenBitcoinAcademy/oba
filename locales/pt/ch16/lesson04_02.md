## Passagem Multisig e Transporte

Em uma configuracao multisig, o PSBT viaja por cada assinante em sequencia ou e distribuido para todos os assinantes em paralelo. Considere um multisig 2-de-3: o coordenador cria o PSBT e o envia para o assinante A e o assinante B simultaneamente. Cada assinante adiciona sua assinatura parcial e retorna o PSBT. O coordenador combina ambos PSBTs, finaliza e transmite.

O mecanismo de transporte nao e especificado pelo BIP 174. PSBTs podem se mover em cartoes SD, como QR codes (usando codificacao UR do BCR-2020-005), via NFC, por e-mail, por servico de compartilhamento de arquivos ou qualquer outro canal. O formato e autocontido. Cada assinante obtem tudo que precisa do proprio PSBT, sem necessidade de canal lateral.

PSBT e texto puro, nao criptografado. Qualquer pessoa que intercepte um PSBT pode ver os valores da transacao, os enderecos envolvidos e as assinaturas parciais coletadas ate o momento. Isso e uma preocupacao de privacidade, nao de seguranca. Um atacante que le um PSBT nao pode roubar fundos porque o PSBT nao contem chaves privadas. Mas o atacante descobre o que o usuario esta gastando e para onde os fundos vao. Para transacoes sensiveis, o PSBT deve ser transportado por um canal criptografado ou em midia fisica que o usuario controla.
