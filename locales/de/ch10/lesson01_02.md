## Node-Typen

Nicht jede Node betreibt dieselbe Software oder hat dieselbe Rolle. Nodes unterscheiden sich in den Daten, die sie speichern, und den Funktionen, die sie bereitstellen.

Eine **Full Node** lädt jeden Block und jede Transaction herunter und validiert sie gegen die Konsensregeln. Sie vertraut niemandem. Sie kann jede Zahlung unabhängig verifizieren. Bitcoin Core ist die verbreitetste Full-Node-Implementierung.

Eine **Mining Node** ist eine Full Node, die auch um die Erstellung neuer Blocks konkurriert. Sie stellt Kandidatenblöcke aus dem Mempool zusammen und führt Proof of Work durch. Mining Nodes erhalten die Block-Subsidy und Transaction Fees, wenn sie einen gültigen Block finden.

Eine **Lightweight Node** (auch SPV-Client genannt) lädt keine vollständigen Blocks herunter. Sie lädt nur Block-Header herunter und fordert den Beweis an, dass bestimmte Transactions existieren. Lightweight Nodes vertrauen Full Nodes in gewissem Maße und tauschen Sicherheit gegen geringeren Ressourcenverbrauch.
