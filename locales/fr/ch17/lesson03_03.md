## Le tweak de securite Taproot

ChillDKG produit une cle publique de groupe, mais cette cle ne peut pas etre utilisee directement comme cle de sortie Taproot. BIP 341 exige que la cle de sortie Q soit une version modifiee d'une cle interne P : Q = P + tagged_hash("TapTweak", P) * G pour une sortie a chemin de cle uniquement.

BIP 445 specifie que ChillDKG applique ce tweak dans le cadre de la generation de cle. Le protocole calcule la cle publique de groupe modifiee et ajuste les parts de cle pour que le groupe puisse signer pour la cle modifiee sans etapes supplementaires pendant la signature. La part de chaque participant est modifiee pour prendre en compte le tweak, et la cle publique de groupe retournee par ChillDKG est la cle de sortie Taproot finale.

Cette integration est necessaire car le tweak doit etre applique de maniere coherente. Si les participants ne sont pas d'accord sur la cle modifiee, la signature echouera. En integrant le tweak dans le protocole DKG lui-meme, ChillDKG garantit que tous les participants derivent la meme cle de sortie et detiennent des parts coherentes avec elle.
