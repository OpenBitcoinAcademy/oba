## Reutilizacao de Nonce e Catastrofica

FROST herda a sensibilidade do Schnorr ao gerenciamento de nonce. Se um assinante usar o mesmo nonce em duas sessoes de assinatura diferentes, um atacante que observa ambos os shares de assinatura pode extrair o share de chave secreta daquele assinante. Com shares extraidos suficientes (t deles), o atacante pode reconstruir a chave secreta completa do grupo e roubar todos os fundos.

Isso nao e uma preocupacao teorica. Geracao deterministica de nonce, que e segura para Schnorr de assinante unico, e perigosa no contexto threshold. Se um coordenador malicioso repetir uma solicitacao de assinatura antiga, um assinante usando nonces deterministicos produziria um novo share de assinatura com o mesmo nonce, vazando seu share. FROST portanto exige nonces criptograficamente aleatorios para cada sessao de assinatura.

Implementacoes devem gerar nonces de uma fonte aleatoria forte e nunca persistir nonces entre sessoes. Se uma sessao de assinatura for abortada, os nonces usados naquela sessao devem ser descartados. Reentrar em uma sessao com os mesmos nonces equivale a reutilizacao de nonce.
