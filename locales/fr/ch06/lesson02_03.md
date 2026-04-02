## Frais de transaction

Les transactions Bitcoin n'ont pas de champ de frais explicite. Les frais sont implicites : c'est la difference entre la valeur totale de toutes les entrees et la valeur totale de toutes les sorties.

Si vos entrees totalisent 100 000 satoshis et vos sorties totalisent 99 800 satoshis, les frais sont de 200 satoshis. Les mineurs collectent ces frais comme incitation pour inclure votre transaction dans un bloc.

Des frais plus eleves signifient une confirmation plus rapide. Quand le reseau est charge, les mineurs privilegient les transactions avec des taux de frais plus eleves. Une transaction qui paie trop peu peut attendre des heures ou des jours.

Le taux de frais depend du poids de la transaction (mesure en octets virtuels, ou vbytes), pas du montant envoye. SegWit a introduit des unites de poids ou les donnees temoin (signatures) sont ristournees par rapport aux autres donnees de transaction. Une transaction envoyant 0.001 BTC peut couter les memes frais qu'une envoyant 1 000 BTC, si les deux ont la meme structure.
