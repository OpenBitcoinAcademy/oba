## Onion Routing

Lightning verpackt Zahlungsanweisungen in Verschlüsselungsschichten, wie eine Zwiebel. Jeder weiterleitende Node entschlüsselt eine Schicht, die nur den nächsten Hop und den weiterzuleitenden Betrag offenbart. Der Node erfährt weder den Absender noch den endgültigen Empfänger noch die Gesamtlänge des Pfads.

Dieses Privatsphäremodell (basierend auf SPHINX) bedeutet, dass Bob, der eine Zahlung von Alice an Carol weiterleitet, nur weiß, dass Alice ihm etwas geschickt hat und dass er es an Carol weiterleiten soll. Er weiß nicht, ob Alice die ursprüngliche Absenderin ist oder ein anderer Weiterleitungs-Node. Er weiß nicht, ob Carol die endgültige Empfängerin ist.

Die Onion hat eine feste Größe unabhängig von der Anzahl der Hops, was eine Pfadlängenanalyse verhindert.
