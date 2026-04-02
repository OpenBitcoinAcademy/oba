## Decrementos de Timelock

Cada salto na cadeia tem um timelock mais curto que o salto anterior. Alice da a Bob 40 blocos. Bob da a Carol 30 blocos. Esse decremento (chamado CLTV delta) garante que cada no de encaminhamento tenha tempo para reivindicar fundos do salto seguinte antes que seu HTLC anterior expire.

Se Carol nao revelar R dentro de 30 blocos, o HTLC de Bob expira e os fundos retornam para Bob. Bob ainda tem 10 blocos para acertar com Alice. Se Bob nao revelar R para Alice dentro de 40 blocos, os fundos de Alice retornam para ela.

O decremento previne um ataque de temporalizacao onde um no seguinte atrasa a revelacao da preimagem ate que o HTLC anterior expire, prendendo o no de encaminhamento.
