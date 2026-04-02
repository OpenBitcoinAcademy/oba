## O Fluxo de Trabalho Somente-Leitura

Um fluxo de assinatura com hardware wallet comeca com uma carteira somente-leitura em um computador conectado a rede. A carteira somente-leitura possui as chaves publicas estendidas (xpubs) de todos os assinantes, mas nenhuma chave privada. Pode gerar enderecos, rastrear saldos e construir transacoes, mas nao pode assinar.

Quando o usuario quer gastar, a carteira somente-leitura cria um PSBT. Ela preenche a transacao nao assinada, anexa dados WITNESS_UTXO para cada input e escreve entradas BIP32_DERIVATION para que a hardware wallet saiba quais chaves usar. O PSBT e exportado para um arquivo em cartao SD ou codificado como QR code animado.

A hardware wallet recebe o PSBT, parseia os inputs e outputs e exibe um resumo para o usuario: quais enderecos recebem fundos, quanto vai para cada um e quanto e pago em taxas. O usuario confirma no dispositivo. A hardware wallet deriva as chaves privadas de sua semente usando os caminhos BIP 32 no PSBT, assina cada input que pode, escreve entradas PARTIAL_SIG e exporta o PSBT atualizado de volta via cartao SD ou QR code.

A carteira somente-leitura importa o PSBT assinado. Se todas as assinaturas necessarias estao presentes, finaliza e extrai a transacao bruta, depois a transmite. Se mais assinaturas sao necessarias (como em uma configuracao multisig), passa o PSBT para o proximo assinante.
