## Encontrando Pares

Quando um no Bitcoin inicia pela primeira vez, ele nao conhece nenhum outro no. Precisa encontrar pelo menos um par para comecar a participar da rede.

O Bitcoin usa **DNS seeds** para descoberta inicial. DNS seeds sao servidores DNS operados por membros da comunidade Bitcoin. Eles retornam enderecos IP de nos completos conhecidos e estaveis. O Bitcoin Core tem varios DNS seeds codificados em seu codigo-fonte.

Se o DNS estiver indisponivel (bloqueado ou censurado), o Bitcoin Core tambem inclui uma lista de enderecos IP codificados como ultimo recurso. Esses enderecos sao atualizados a cada lancamento de software.

Uma vez que um no se conecta ao seu primeiro par, pode pedir a ele enderecos adicionais. O no constroi um conjunto diversificado de conexoes, tipicamente 8 de saida e ate 125 de entrada.
