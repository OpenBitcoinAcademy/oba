## Bedingte Zahlungen über mehrere Hops

Hash Time-Locked Contracts (HTLCs) ermöglichen Zahlungen zwischen Parteien, die keinen direkten Channel teilen. Ein HTLC kombiniert zwei Bedingungen: einen Hashlock (ein Secret offenlegen, um Mittel zu beanspruchen) und einen Timelock (wenn das Secret nicht rechtzeitig offengelegt wird, gehen die Mittel an den Absender zurück).

Carol erzeugt ein zufälliges Secret R und sendet den Hash H(R) an Alice in einer Payment Invoice. Alice kennt R nicht. Sie erstellt einen HTLC mit Bob: "Zahle 1,001 BTC, wenn du das Preimage von H(R) innerhalb von 40 Blocks offenlegst." Bob erstellt einen HTLC mit Carol: "Zahle 1,000 BTC, wenn du das Preimage von H(R) innerhalb von 30 Blocks offenlegst."
