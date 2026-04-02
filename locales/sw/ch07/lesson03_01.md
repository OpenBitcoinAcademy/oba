## SegWit Asilia (P2WPKH na P2WSH)

Segregated Witness ilihamisha data ya saini nje ya hati ya ingizo na kuiweka katika muundo tofauti wa ushahidi. Hii ilirekebisha ubadilishaji wa muamala (watu wa tatu wangeweza kubadilisha txid kwa kurekebisha usimbaji wa saini) na kuwezesha punguzo la ushahidi kwa ada.

P2WPKH (Pay to Witness Public Key Hash) ni sawa na P2PKH ya SegWit. Hati ya matokeo ina: OP_0 <hash ya baiti 20 ya funguo ya umma>. Ushahidi unatoa saini na funguo ya umma. Hati ya ingizo ni tupu.

P2WSH (Pay to Witness Script Hash) ni sawa na P2SH ya SegWit. Hati ya matokeo ina: OP_0 <hash ya baiti 32 ya hati>. Ushahidi unatoa hati na data inayohitajika. P2WSH inatumia hash ya SHA-256 ya baiti 32 badala ya HASH160 ya baiti 20 ya P2SH, na kutoa upinzani bora wa mgongano.
