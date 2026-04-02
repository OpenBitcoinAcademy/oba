## Scripts de Desbloqueio

Para gastar um output bloqueado, o gastador fornece um script de desbloqueio, chamado scriptSig. Para P2PKH, o script de desbloqueio contem uma assinatura digital e a chave publica do gastador.

O Bitcoin executa esses scripts em uma maquina de pilha. Primeiro, o script de desbloqueio roda e coloca dados na pilha. Depois, a pilha e copiada, e o script de bloqueio roda contra essa pilha copiada. Os dois scripts nunca se combinam em um so. Essa separacao foi introduzida em 2010 para corrigir uma vulnerabilidade de seguranca.

Se o script de bloqueio terminar com um valor verdadeiro no topo da pilha, o gasto e valido. Se a pilha estiver vazia ou o valor do topo for zero, o gasto falha.
