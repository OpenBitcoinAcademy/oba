## BitVM : verification de calcul

BitVM permet la verification de calculs Turing-complets sur Bitcoin sans changement de consensus. Un prouveur affirme qu'une fonction f(x) = y. Si l'affirmation est correcte, rien ne se passe sur la chaine. Si elle est incorrecte, n'importe quel verificateur peut soumettre une preuve de fraude et le prouveur perd une caution.

BitVM2 (2024-2025) a rendu la verification sans permission : n'importe qui peut contester une fausse affirmation, et les litiges se resolvent en trois transactions sur la chaine. Le systeme implemente un verificateur SNARK dans Bitcoin Script, permettant une verification efficace de calculs complexes.

Les applications incluent les ponts sans confiance (deplacer des bitcoins entre les chaines), les rollups (regrouper des transactions hors chaine avec verification sur la chaine) et tout protocole qui beneficie de "verifier, ne pas calculer" sur Bitcoin.
