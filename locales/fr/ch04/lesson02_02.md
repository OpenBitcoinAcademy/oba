## De la cle privee a la cle publique

A partir d'une cle privee, vous pouvez calculer une cle publique. Cela utilise la multiplication sur courbe elliptique, une operation mathematique facile a realiser dans un sens et impossible a inverser par calcul.

La cle privee est un nombre. Multipliez-le par un point specifique sur une courbe connue (appele le point generateur G), et vous obtenez un autre point sur la courbe. Ce resultat est votre cle publique.

A partir de la seule cle publique, personne ne peut determiner la cle privee. N'importe qui peut verifier une signature numerique par rapport a une cle publique, confirmant que le signataire controle la cle privee correspondante, sans que la cle privee soit revelee. Cette relation a sens unique est le fondement de la securite de Bitcoin.
