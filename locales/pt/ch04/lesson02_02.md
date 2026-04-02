## Da Chave Privada a Chave Publica

A partir de uma chave privada, voce pode calcular uma chave publica. Isso usa multiplicacao de curva eliptica, uma operacao matematica facil de realizar em uma direcao e computacionalmente inviavel de reverter.

A chave privada e um numero. Multiplique-o por um ponto especifico em uma curva conhecida (chamado ponto gerador G) e voce obtem outro ponto na curva. Esse resultado e sua chave publica.

Dada apenas a chave publica, ninguem pode determinar a chave privada. Qualquer pessoa pode verificar uma assinatura digital contra uma chave publica, confirmando que o assinante controla a chave privada correspondente, sem que a chave privada seja revelada. Essa relacao unidirecional e a base da seguranca do Bitcoin.
