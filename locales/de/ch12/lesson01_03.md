## Difficulty-Anpassung

Bitcoin passt seine Difficulty alle 2.016 Blocks an, ungefähr alle zwei Wochen. Das Ziel ist ein durchschnittliches Block-Intervall von 10 Minuten.

An jedem Anpassungspunkt berechnen Nodes, wie lange die vorherigen 2.016 Blocks gedauert haben. Wenn Blocks schneller als 10 Minuten im Durchschnitt kamen, sinkt das Target (die Difficulty steigt). Wenn Blocks langsamer kamen, steigt das Target (die Difficulty sinkt). Die Formel vergleicht die tatsächlich vergangene Zeit mit den erwarteten 20.160 Minuten.

Eine Sicherheitsvorkehrung verhindert, dass sich die Difficulty in einer einzelnen Anpassung um mehr als den Faktor vier ändert. Das begrenzt, wie schnell die Difficulty steigen oder fallen kann. Der Mechanismus ist vollständig algorithmisch. Kein Gremium stimmt darüber ab. Jede Node berechnet dasselbe neue Target aus denselben Chain-Daten.
