## Roteamento Baseado na Origem

O remetente, nao a rede, escolhe o caminho de pagamento. Cada no Lightning mantem uma visao local da topologia da rede (quais nos existem, quais canais os conectam, suas capacidades e taxas). O remetente calcula uma rota a partir dessas informacoes.

Isso difere do roteamento de internet, onde cada roteador decide independentemente o proximo salto. No Lightning, o remetente tem controle total sobre o caminho. Os nos intermediarios seguem instrucoes de encaminhamento sem conhecer a rota completa.
