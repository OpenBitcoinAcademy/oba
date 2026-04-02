## Was Output Descriptors sind

Ein Output Descriptor ist ein String, der vollständig beschreibt, wie eine Menge von Bitcoin Addresses abgeleitet wird. BIPs 380 bis 389 definieren die Descriptor-Sprache. Jeder Descriptor gibt den Script-Typ, die beteiligten Keys und die Ableitungspfade an.

Ein Descriptor wie `wpkh([d34db33f/84h/0h/0h]xpub.../0/*)` sagt einem Wallet alles Nötige: verwende P2WPKH, leite von diesem Extended Public Key an diesem Pfad ab und erzeuge sequenzielle Addresses. Das Wallet muss den Script-Typ oder das Ableitungsschema nicht erraten. Der Descriptor ist in sich geschlossen.

Vor Descriptors verließen sich Wallets auf Konventionen. BIP 44 legte fest, dass HD Wallets einen bestimmten Ableitungspfad verwenden sollen. BIP 49 fügte einen anderen Pfad für P2SH-SegWit hinzu. BIP 84 einen weiteren für natives SegWit. Ein Wallet, das einen xpub importierte, musste alle Konventionen durchprobieren und hoffen, richtig zu raten. Descriptors ersetzen Rätselraten durch explizite Deklarationen.
