## Le mecanisme de revocation

Quand Alice et Bob mettent a jour le solde de leur canal, ils invalident l'engagement precedent en echangeant des secrets de revocation. Chaque partie detient maintenant un secret qui peut punir l'autre pour la diffusion de l'ancien etat.

Si Bob diffuse un ancien engagement ou il avait 0.5 BTC (mais le solde actuel ne lui donne que 0.3 BTC), Alice peut utiliser le secret de revocation de Bob pour reclamer l'integralite du solde du canal. Bob perd tout.

Cette penalite rend la triche economiquement irrationnelle. Le cout de la tentative de fraude (perdre tous les fonds) depasse de loin le gain potentiel (reclamer un ancien solde legerement plus eleve). Le comportement honnete est la strategie dominante.
