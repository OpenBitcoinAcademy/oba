## Construindo Sistemas Bitcoin com Seguranca

Desenvolvedores construindo sobre Bitcoin enfrentam um desafio unico: bugs podem perder dinheiro. Uma falha na geracao de chaves, construcao de transacao ou validacao de assinatura pode resultar em fundos sendo permanentemente roubados ou perdidos.

O modelo de consenso descentralizado significa que nao ha autoridade para reverter uma transacao equivocada. Codigo que lida com chaves privadas deve trata-las como os dados mais sensiveis do sistema. Chaves devem ser geradas de fontes de entropia de alta qualidade, armazenadas criptografadas em repouso e apagadas da memoria apos o uso.

Cada componente que toca chaves ou constroi transacoes deve ser auditado, testado contra vetores conhecidos e submetido a revisao adversarial.
