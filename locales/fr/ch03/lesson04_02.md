## Le paysage des implementations

Plusieurs equipes maintiennent des implementations Bitcoin independantes dans differents langages de programmation.

**btcd** (Go) est une implementation de noeud complet ecrite a partir de zero. Elle alimente l'infrastructure de plusieurs entreprises Bitcoin et sert de base au lnd du Lightning Network.

**bcoin** (JavaScript) est une implementation de noeud complet modulaire concue pour l'ecosysteme Node.js, avec un portefeuille integre et une API HTTP.

**Bitcoin Knots** est un fork de Bitcoin Core maintenu par Luke Dashjr. Il inclut des options de configuration supplementaires et des politiques par defaut plus strictes.

**rust-bitcoin** est une bibliotheque pour travailler avec les structures de donnees Bitcoin en Rust. Elle fournit des outils de serialisation, d'analyse et de scripting sans executer de noeud complet.

**libbitcoin** (C++) est une boite a outils independante pour construire des applications Bitcoin, incluant une implementation de noeud complet appelee libbitcoin-node.

D'autres implementations existent en Python, Java, Scala et C#. Chacune amene une communaute de developpeurs differente dans l'ecosysteme Bitcoin.
