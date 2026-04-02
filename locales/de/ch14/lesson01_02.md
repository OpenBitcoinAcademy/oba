## Warum Tweaking funktioniert

Der Tweak bindet den Script-Baum kryptographisch an den Public Key. Eine Änderung an einem Script im Baum ändert den Merkle Root, der ändert den Tweak, der ändert Q. Ein Script, das bei der Output-Erstellung in Q committet wurde, kann nachträglich nicht verändert werden.

Der Private Key wird identisch getweakt: tweaked_privkey = (privkey + t) mod n. Der Inhaber des internen Private Key kann den getweakten Private Key berechnen und direkt signieren. Das ist der Key Path Spend.

Tagged Hashes verhindern Kollisionen zwischen verschiedenen Verwendungen der Hash-Funktion. Der Tag "TapTweak" ist domänengetrennt: SHA256(SHA256("TapTweak") || SHA256("TapTweak") || data). Verschiedene Tags erzeugen bei gleichen Daten unterschiedliche Hash-Ausgaben und eliminieren protokollübergreifende Angriffe.
