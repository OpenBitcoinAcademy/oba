Que se passe-t-il si un signataire FROST reutilise le meme nonce dans deux sessions de signature differentes ?

- La signature finale devient invalide mais aucune donnee secrete n'est divulguee
- Un attaquant peut extraire la part de cle secrete de ce signataire a partir des deux parts de signature
- Le coordinateur detecte la reutilisation et interrompt les deux sessions
