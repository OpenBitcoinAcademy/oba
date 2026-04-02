## O Fechamento por Penalidade

Se uma parte transmitir uma transacao de commitment revogada (antiga), a outra parte pode usar o segredo de revogacao para reivindicar todo o saldo do canal. O trapaceiro perde todos os fundos no canal.

Esse e o mecanismo de aplicacao economica. Transmitir estado antigo e detectavel (a outra parte monitora o blockchain em busca de commitments revogados) e punivel (perda total de fundos). A penalidade torna canais Lightning confiaveis sem exigir confianca entre as partes.

A parte honesta deve estar online (ou ter um servico de watchtower monitorando em seu nome) para detectar e responder a uma transmissao de commitment revogado dentro da janela de timelock.
