## Erstellen und Verifizieren

Das Erstellen einer Signature erfordert zwei Eingaben: den Private Key und die Nachricht (die Transaction-Daten oder einen Hash davon). Der Signature-Algorithmus erzeugt einen Signature-Wert, der sowohl für den Key als auch für die Nachricht einzigartig ist.

Das Verifizieren einer Signature erfordert drei Eingaben: den Public Key, die Nachricht und die Signature. Der Verifikationsalgorithmus gibt wahr oder falsch aus. Bei wahr wurde die Signature vom Inhaber des zugehörigen Private Key für genau diese Nachricht erzeugt.

Der Private Key wird beim Erstellen oder Verifizieren nie offengelegt. Jeder kann verifizieren, aber nur der Key-Inhaber kann signieren.
