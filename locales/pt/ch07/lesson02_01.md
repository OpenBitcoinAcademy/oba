## Pay to Script Hash (P2SH)

O P2SH separa o script do endereco. Em vez de codificar o script de bloqueio completo no output, o P2SH codifica apenas um hash do script. O gastador revela o script completo (chamado redeem script) ao gastar.

O script de output contem: OP_HASH160 <script_hash> OP_EQUAL. Para gastar, o input fornece o redeem script e quaisquer dados que o redeem script exija (assinaturas, chaves publicas). A rede faz hash do redeem script fornecido e verifica se corresponde ao hash no output.

Enderecos P2SH comecam com "3" na mainnet. Eles habilitam condicoes de gasto complexas (multisig, timelocks) sem exigir que o remetente conheca os detalhes. O remetente paga para um hash curto. O destinatario lida com a complexidade.
