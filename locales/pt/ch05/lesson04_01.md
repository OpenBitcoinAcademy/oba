## Frases Semente BIP39

O BIP39 codifica uma semente de carteira como uma sequencia de 12 ou 24 palavras em ingles, chamada frase de recuperacao (ou frase semente). Cada palavra vem de uma lista padronizada de 2.048 palavras.

A codificacao em palavras serve a dois propositos. Palavras sao mais faceis de escrever, ler e transcrever do que hexadecimal bruto. E a ultima palavra inclui um checksum que detecta erros de transcricao.

Uma frase de 12 palavras codifica 128 bits de entropia. Uma frase de 24 palavras codifica 256 bits. Ambas sao fortes o suficiente para as necessidades de seguranca atuais. A frase e convertida em uma semente de 512 bits usando PBKDF2 com 2.048 rodadas de HMAC-SHA512.
