## Ouvrir un canal

Ouvrir un canal necessite une transaction sur la chaine : la transaction de financement. Elle cree une sortie multisig 2-of-2. Avant la diffusion, les deux parties signent la premiere transaction d'engagement (le filet de securite qui restitue les fonds si le canal echoue).

Le canal est ouvert une fois la transaction de financement confirmee. Les deux parties peuvent commencer a mettre a jour le solde en echangeant de nouvelles transactions d'engagement. Le cout sur la chaine est un seul frais de transaction pour la transaction de financement.
