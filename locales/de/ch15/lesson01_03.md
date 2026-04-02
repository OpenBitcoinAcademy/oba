## Was Miniscript löst

Miniscript, spezifiziert in BIP 379, ist eine strukturierte Darstellung von Bitcoin Script. Es bildet eine definierte Menge komponierbarer Fragmente auf Script-Opcodes ab. Jedes Fragment hat bekannte Eigenschaften: seinen Typ, seine Ressourcenkosten und die benötigten Witness-Daten.

Weil die Abbildung strukturiert ist, kann Software einen Miniscript-Ausdruck analysieren und jeden Ausgabepfad, jeden benötigten Key, jede Timelock-Beschränkung und die exakte Witness-Größe für jeden Pfad bestimmen. Zwei Ausdrücke, die dieselbe Policy kodieren, können verglichen und als äquivalent nachgewiesen werden.

Komposition wird sicher. Ein Wallet kann zwei Miniscript-Fragmente nehmen, sie mit AND oder OR kombinieren und den Ressourcenverbrauch des resultierenden Scripts vor dem Senden berechnen. Überschreitet ein Ausführungspfad die Konsens-Limits, lehnt der Compiler die Komposition zur Build-Zeit ab, nicht zur Ausgabezeit.
