## Public-Child-Key-Ableitung

Eine Eigenschaft von HD Wallets ist die Möglichkeit, öffentliche Child Keys aus einem übergeordneten Public Key abzuleiten, ohne den übergeordneten Private Key zu kennen.

Ein Webserver, der Empfangs-Addresses für einen Onlineshop erzeugt, kann nur den übergeordneten Public Key halten. Er leitet für jeden Kunden neue Empfangs-Addresses ab, ohne Zugang zu irgendeinem Private Key zu haben. Die Private Keys, die zum Ausgeben nötig sind, existieren nur auf einem separaten, gesicherten Gerät.

Diese Trennung von öffentlichem und privatem Key-Raum ist ein großer Sicherheitsvorteil. Das Online-System, das kundenseitige Operationen abwickelt, berührt nie die Ausgabe-Keys.
