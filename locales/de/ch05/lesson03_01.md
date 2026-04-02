## BIP32: Der HD-Wallet-Standard

BIP32 definiert, wie ein Baum von Keys aus einem einzelnen Seed abgeleitet wird. Der Prozess verwendet HMAC-SHA512, um bei jedem Ableitungsschritt einen Child Key und einen Chain Code zu erzeugen. Der Chain Code fließt in die nächste Ableitung ein und verhindert, dass jemand Geschwister-Keys berechnen kann.

Der Master Key sitzt an der Wurzel des Baums. Aus ihm wird eine Folge von Child Keys abgeleitet, jeder durch eine Indexnummer identifiziert (0, 1, 2, ...). Jeder Child kann eigene Children erzeugen und so einen Baum beliebiger Tiefe bilden.

Diese Baumstruktur gibt Wallets organisatorische Stärke. Verschiedene Zweige des Baums dienen verschiedenen Zwecken, und der gesamte Baum lässt sich aus dem ursprünglichen Seed wiederherstellen.
