## A Frase Secreta Opcional

O BIP39 suporta uma frase secreta opcional que e misturada na derivacao da semente. As mesmas 12 palavras com frases secretas diferentes produzem sementes diferentes e carteiras diferentes.

Isso fornece negacao plausivel: um usuario pode criar uma carteira isca com uma frase secreta e uma carteira real com outra. Sob coercao, ele pode revelar a frase secreta da isca sem expor os fundos principais.

O risco: esquecer a frase secreta significa perder acesso a carteira. Nao ha mecanismo de recuperacao. A frase secreta nao e armazenada em lugar nenhum. Ela existe apenas na memoria do usuario ou em um backup fisico.
