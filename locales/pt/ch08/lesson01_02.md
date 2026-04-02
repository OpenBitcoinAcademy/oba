## Criando e Verificando

Criar uma assinatura requer duas entradas: a chave privada e a mensagem (os dados da transacao, ou um hash deles). O algoritmo de assinatura produz um valor de assinatura que e unico para a chave e a mensagem.

Verificar uma assinatura requer tres entradas: a chave publica, a mensagem e a assinatura. O algoritmo de verificacao retorna verdadeiro ou falso. Se verdadeiro, a assinatura foi produzida pelo detentor da chave privada correspondente para aquela mensagem exata.

A chave privada nunca e revelada durante a criacao ou verificacao. Qualquer pessoa pode verificar, mas so o detentor da chave pode assinar.
