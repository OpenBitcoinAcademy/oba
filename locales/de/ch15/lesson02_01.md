## Drei Schichten: Policy, Miniscript, Script

Miniscript arbeitet in einer Drei-Schichten-Architektur. Oben steht eine menschenlesbare Policy-Sprache. In der Mitte ist Miniscript selbst. Unten liegt Bitcoin Script.

Eine Policy beschreibt, was du willst: "Alice UND Bob, ODER Carol nach 90 Tagen." Du schreibst das als `or(and(pk(Alice),pk(Bob)),and(pk(Carol),older(12960)))`. Die Policy-Sprache ist für Menschen. Sie benennt Keys und Timelocks ohne Rücksicht auf Opcodes.

Ein Compiler übersetzt die Policy in einen Miniscript-Ausdruck: `or_d(and_v(v:pk(Alice),pk(Bob)),and_v(v:pk(Carol),older(12960)))`. Der Miniscript-Ausdruck spezifiziert exakt, wie die Bedingungen komponiert werden, einschließlich der verwendeten Fragmenttypen an jeder Position. Von dort bildet der Ausdruck direkt auf Bitcoin Script Opcodes ab. Jedes Miniscript-Fragment hat genau eine Script-Kodierung.
