## Der ECDSA-Algorithmus

ECDSA (Elliptic Curve Digital Signature Algorithm) war Bitcoins ursprüngliches Signature-Verfahren. Es arbeitet auf der elliptischen Kurve secp256k1, derselben Kurve, die für die Key-Erzeugung verwendet wird.

Zum Signieren einer Nachricht nimmt der Algorithmus den Private Key und einen Hash der Nachricht, kombiniert sie mit einer Zufallszahl (genannt k, oder Nonce) und erzeugt zwei Werte: r und s. Zusammen bilden (r, s) die Signature.

Zur Verifikation nimmt der Algorithmus den Public Key, den Nachrichten-Hash und die Signature (r, s). Er führt Berechnungen auf der elliptischen Kurve durch, um zu prüfen, ob die Signature mit dem Public Key und der Nachricht konsistent ist. Kein Private Key wird für die Verifikation benötigt.
