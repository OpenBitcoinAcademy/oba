## Bijektive Abbildung

Die Beziehung zwischen Miniscript und Script ist bijektiv für die unterstützte Fragmentmenge. Jeder Miniscript-Ausdruck bildet auf genau eine Script-Kodierung ab, und jedes unterstützte Script-Muster bildet auf genau einen Miniscript-Ausdruck zurück.

Betrachte das Fragment `pk(K)`: es kodiert zu `<K> OP_CHECKSIG`. Das Fragment `and_v(v:pk(A),pk(B))` kodiert zu `<A> OP_CHECKSIGVERIFY <B> OP_CHECKSIG`. Es gibt keine Mehrdeutigkeit. Aus dem Script kannst du das ursprüngliche Miniscript rekonstruieren.

Diese Eigenschaft macht Roundtrip-Analyse möglich. Ein Wallet kann ein Script von einer anderen Partei empfangen, es zu Miniscript dekodieren und die Ausgabebedingungen bestimmen, ohne der Beschreibung des Absenders zu vertrauen. Zwei Wallets verschiedener Teams können sich auf eine Policy einigen, unabhängig zu Miniscript kompilieren und verifizieren, dass sie dasselbe Script erzeugt haben.

Nicht jedes gültige Bitcoin Script lässt sich in Miniscript darstellen. Die Fragmentmenge deckt die Muster ab, die für praktische Ausgabepolicies nötig sind: Keys, Hashes, Timelocks, Schwellenwerte und boolesche Kombinationen. Scripts, die ungewöhnliche Opcode-Sequenzen oder Stack-Manipulation verwenden, liegen außerhalb der Miniscript-Teilmenge.
