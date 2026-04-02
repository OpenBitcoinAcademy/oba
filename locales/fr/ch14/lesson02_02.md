## MuSig2 comme cle interne

La cle interne P peut etre un agregat MuSig2 de plusieurs cles publiques. Si Alice et Bob aggregent leurs cles en P en utilisant MuSig2, la cle modifiee Q engage sur les deux. Quand ils cooperent, ils produisent une seule signature Schnorr sur Q.

Sur la chaine, ce multisig 2-of-2 semble identique a une transaction a signataire unique. La sortie fait 34 octets. Le temoin fait 64 octets. Aucun observateur ne peut determiner que deux parties etaient impliquees.

BitGo a rapporte qu'une entree par chemin de cle MuSig2 coute 57,5 octets virtuels, contre 104,5 vbytes pour une entree multisig P2WSH SegWit native. Les economies proviennent de l'elimination des cles publiques et signatures par signataire que les scripts multisig exigent sur la chaine.
