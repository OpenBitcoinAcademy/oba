## Forks e Reorganizacoes

Dois mineradores as vezes encontram blocos validos quase ao mesmo tempo. Quando isso acontece, partes da rede veem um bloco primeiro, e outras partes veem o outro. A cadeia se bifurca temporariamente em dois ramos concorrentes.

Isso e um fork natural, nao um ataque. Cada no trabalha no ramo que recebeu primeiro. O empate se resolve quando o proximo bloco e encontrado. Se um minerador estende um ramo, esse ramo agora tem mais proof of work cumulativo. Nos no ramo mais curto mudam para o mais longo. Essa troca e chamada reorganizacao. As transacoes no bloco abandonado retornam ao mempool para inclusao em um bloco futuro.

A maioria dos forks naturais se resolve em um bloco. Reorganizacoes mais profundas sao raras e se tornam exponencialmente menos provaveis com cada confirmacao adicional.
