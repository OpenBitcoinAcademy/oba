## Transaction Pinning

Transaction pinning e um ataque onde uma parte maliciosa cria uma transacao filha de taxa baixa que torna caro ou impossivel para a parte honesta aumentar a taxa da mae.

Em um protocolo de duas partes (como um canal da Lightning Network), uma parte pode transmitir um descendente grande e de taxa baixa de uma transacao compartilhada. A tentativa de CPFP da outra parte precisaria pagar pelo descendente grande do atacante, tornando o aumento de taxa proibitivamente caro.

Anchor outputs sao uma contramedida. Cada parte recebe um output pequeno na transacao compartilhada que pode gastar com CPFP. Regras limitam quantos descendentes cada anchor pode ter, prevenindo o ataque de pinning.
