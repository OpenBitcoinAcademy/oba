## Fee Sniping

Le fee sniping est une attaque theorique ou un mineur, au lieu de construire sur le dernier bloc, re-mine un bloc precedent pour voler ses frais de transaction. Si les frais dans le bloc N sont assez eleves, un mineur pourrait trouver plus rentable de forker la chaine au bloc N-1 et collecter ces frais lui-meme.

Bitcoin Core se defend contre cela en reglant le locktime des nouvelles transactions a la hauteur de bloc actuelle. Une transaction verrouillee au bloc 800 000 ne peut pas etre incluse dans le bloc 799 999. Cela rend moins rentable de re-miner les anciens blocs, car les nouvelles transactions creees depuis seraient indisponibles.

Le fee sniping ne s'est pas produit sur le reseau Bitcoin. La defense est preventive.
