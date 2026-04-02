## Funcoes, Nao Pessoas

O fluxo de trabalho PSBT define sete funcoes: Creator, Updater, Signer, Combiner, Finalizer, Extractor e (no PSBTv2) Constructor. Cada funcao executa uma etapa. Uma unica pessoa ou programa pode preencher multiplas funcoes, mas a separacao importa porque cada funcao tem requisitos de confianca diferentes e acesso diferente a informacoes.

O Creator constroi o PSBT inicial. No v0 (BIP 174), o Creator produz a transacao nao assinada e a envolve em um PSBT com mapas de input e output vazios. No v2 (BIP 370), o Creator define campos globais como a versao da transacao e locktime, mas ainda nao inclui inputs ou outputs. Essa tarefa cabe ao Constructor.

A funcao Constructor existe apenas no PSBTv2. Ele adiciona inputs e outputs ao PSBT. Multiplos Constructors podem colaborar: um adiciona os inputs que controla, outro adiciona seus inputs, e cada um adiciona os outputs que precisa. Isso habilita construcao interativa de transacao onde nenhuma parte conhece a forma completa da transacao ate que todos os Constructors tenham contribuido.
