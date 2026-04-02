## O Mecanismo de Revogacao

Quando Alice e Bob atualizam o saldo do canal, eles invalidam o commitment anterior trocando segredos de revogacao. Cada parte agora detem um segredo que pode punir a outra por transmitir o estado antigo.

Se Bob transmitir um commitment antigo onde ele tinha 0,5 BTC (mas o saldo atual lhe da apenas 0,3 BTC), Alice pode usar o segredo de revogacao de Bob para reivindicar todo o saldo do canal. Bob perde tudo.

Essa penalidade torna a trapacea economicamente irracional. O custo de tentar fraude (perder todos os fundos) excede em muito o ganho potencial (reivindicar um saldo antigo ligeiramente maior). Comportamento honesto e a estrategia dominante.
