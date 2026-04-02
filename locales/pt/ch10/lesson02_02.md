## O Handshake

Dois nos estabelecem uma conexao com um handshake version/verack sobre TCP na porta 8333 (a porta padrao da rede Bitcoin).

O no que se conecta envia uma mensagem `version` contendo sua versao de protocolo, altura de bloco, hora atual e os servicos que suporta. O no receptor responde com sua propria mensagem `version`. Cada no entao envia um `verack` (confirmacao de versao) para confirmar.

Apos o handshake ser completado, os nos podem trocar dados. Se as versoes de protocolo forem incompativeis, a conexao e encerrada.

Nos compartilham enderecos de pares conhecidos usando mensagens `addr` e `getaddr`. Quando um no aprende novos enderecos, os armazena e pode se conectar a eles depois. Esse protocolo de gossip permite que a rede cresca e se recupere sem nenhum diretorio central.
