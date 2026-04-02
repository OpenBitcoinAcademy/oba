## Quando o Key Path Falha

Se o key path nao pode ser usado (um assinante esta indisponivel, um timelock precisa ser exercido), o script path fornece uma alternativa. O gastador revela um script da arvore comprometida em Q, prova que foi comprometido no momento da criacao e satisfaz aquele script.

A witness para um gasto por script path contem: os dados que satisfazem o script (assinaturas, preimagens), o proprio script e um control block. O control block contem a chave interna P, um byte de versao de folha e a prova Merkle (os hashes irmaos ao longo do caminho da folha do script ate a raiz).

O verificador reconstroi a raiz Merkle a partir do script revelado e da prova, calcula o tweak esperado e verifica que Q e igual a P + tweak * G. Se a matematica confere, o script foi comprometido no momento da criacao.
