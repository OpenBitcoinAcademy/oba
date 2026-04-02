## P2PKH Passo a Passo

Veja o que acontece quando um script P2PKH e executado:

1. A assinatura e a chave publica sao colocadas na pilha (do script de desbloqueio).
2. OP_DUP duplica a chave publica na pilha.
3. OP_HASH160 aplica hash no elemento do topo (a chave publica duplicada) com SHA-256 e depois RIPEMD-160.
4. O hash de chave publica esperado e colocado na pilha (do script de bloqueio).
5. OP_EQUALVERIFY remove os dois elementos do topo e verifica se sao iguais. Se diferirem, o script falha.
6. OP_CHECKSIG remove a assinatura e a chave publica, verifica a assinatura contra os dados da transacao. Se valida, coloca true.

O resultado: true na pilha. O gasto esta autorizado. A assinatura prova que o gastador controla a chave privada, sem revela-la.
