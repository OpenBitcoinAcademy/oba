## Qui paie les frais de transaction ?

L'expediteur paie les frais de transaction. Les frais ne sont pas un paiement separe. C'est la difference entre la valeur totale des entrees et la valeur totale des sorties.

Si Alice cree une transaction avec des entrees totalisant 100 000 satoshis et des sorties totalisant 99 500 satoshis, les 500 satoshis restants constituent les frais. Le mineur qui inclut la transaction dans un bloc les collecte.

Il n'y a pas de regle imposant des frais minimum. Mais les mineurs privilegient les transactions avec des taux de frais plus eleves par unite de poids. Une transaction avec des frais trop bas peut attendre dans le mempool pendant des heures, des jours, ou ne jamais etre confirmee.
