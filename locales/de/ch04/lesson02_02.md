## Vom Private Key zum Public Key

Aus einem Private Key kannst du einen Public Key berechnen. Das nutzt elliptische Kurvenmultiplikation, eine mathematische Operation, die in eine Richtung leicht durchführbar und in die Gegenrichtung rechnerisch nicht umkehrbar ist.

Der Private Key ist eine Zahl. Multipliziere sie mit einem bestimmten Punkt auf einer bekannten Kurve (dem Generatorpunkt G) und du erhältst einen anderen Punkt auf der Kurve. Dieses Ergebnis ist dein Public Key.

Aus dem Public Key allein kann niemand den Private Key bestimmen. Jeder kann eine digitale Signature gegen einen Public Key prüfen und so bestätigen, dass der Signierer den zugehörigen Private Key kontrolliert, ohne dass der Private Key offengelegt wird. Diese Einwegbeziehung ist die Grundlage der Sicherheit von Bitcoin.
