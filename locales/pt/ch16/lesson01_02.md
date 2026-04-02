## O que e um PSBT

Um Partially Signed Bitcoin Transaction (PSBT) e um formato binario definido pelo BIP 174. Ele envolve uma transacao nao assinada junto com os metadados que cada participante precisa para fazer seu trabalho. Criadores anexam dados de UTXO. Atualizadores adicionam caminhos de derivacao. Assinantes adicionam assinaturas. O formato carrega tudo que e necessario para que cada funcao atue de forma independente.

O binario comeca com cinco bytes magicos: `0x70736274FF`. Os primeiros quatro bytes soletram "psbt" em ASCII. O quinto byte, `0xFF`, marca o separador. Qualquer ferramenta que receba um blob pode verificar esses cinco bytes para confirmar que e um PSBT antes de parsear.

Apos o magic vem uma sequencia de mapas chave-valor. Cada entrada de mapa tem um tipo de chave, um payload de chave e um payload de valor. Um byte zero (0x00) termina cada mapa. O primeiro mapa e o mapa global. Mapas subsequentes alternam entre mapas por input e por output, um para cada input e output na transacao.
