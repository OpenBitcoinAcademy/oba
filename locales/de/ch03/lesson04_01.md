## Warum mehrere Implementierungen wichtig sind

Ein Protokoll, das durch eine einzige Implementierung definiert wird, ist fragil. Wenn Bitcoin Core einen Bug hat, ist jede Node, die es ausführt, gleichzeitig betroffen. Ein Bug im Konsenscode könnte das Netzwerk spalten oder ungültige Transactions zulassen.

Mehrere unabhängige Implementierungen desselben Protokolls stärken das Ökosystem. Wenn zwei Implementierungen sich uneinig sind, ob ein Block gültig ist, offenbart die Uneinigkeit einen Bug in einer von ihnen. Das geschah 2013, als ein Unterschied zwischen BerkeleyDB und LevelDB einen unbeabsichtigten Chain-Split verursachte. Der Vorfall zeigte sowohl das Risiko einer Monokultur als auch den Wert von Vielfalt.

Ein gesundes Protokoll lässt sich von mehreren Teams anhand der Spezifikation implementieren. Je mehr Implementierungen sich in jedem Grenzfall einig sind, desto größer ist das Vertrauen des Netzwerks, dass die Regeln korrekt und wohldefiniert sind.
