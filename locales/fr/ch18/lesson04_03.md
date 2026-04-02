## La fermeture avec penalite

Si une partie diffuse une transaction d'engagement revoquee (ancienne), l'autre partie peut utiliser le secret de revocation pour reclamer l'integralite du solde du canal. Le tricheur perd tous ses fonds dans le canal.

C'est le mecanisme d'application economique. Diffuser un ancien etat est detectable (l'autre partie surveille la blockchain pour les engagements revoques) et punissable (perte totale des fonds). La penalite rend les canaux Lightning fiables sans necessiter de confiance entre les parties.

La partie honnete doit etre en ligne (ou avoir un service de tour de guet surveillant en son nom) pour detecter et repondre a la diffusion d'un engagement revoque dans la fenetre du verrou temporel.
