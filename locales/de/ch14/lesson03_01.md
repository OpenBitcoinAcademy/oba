## Wenn der Key Path fehlschlägt

Wenn der Key Path nicht genutzt werden kann (ein Signer ist nicht verfügbar, ein Timelock muss ausgeübt werden), bietet der Script Path eine Alternative. Der Absender legt ein Script aus dem in Q committeten Baum offen, beweist, dass es bei der Erstellung committet wurde, und erfüllt dieses Script.

Die Witness für einen Script Path Spend enthält: die Daten, die das Script erfüllen (Signatures, Preimages), das Script selbst und einen Control Block. Der Control Block enthält den Internal Key P, ein Leaf-Version-Byte und den Merkle Proof (die Geschwister-Hashes entlang des Pfades vom Script Leaf zur Root).

Der Verifizierer rekonstruiert den Merkle Root aus dem offengelegten Script und dem Proof, berechnet den erwarteten Tweak und prüft, ob Q gleich P + Tweak * G ist. Stimmt die Mathematik, wurde das Script bei der Erstellung committet.
