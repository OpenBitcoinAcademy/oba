## Was ist ein Merkle Tree?

Ein Merkle Tree ist ein binärer Baum aus Hashes. Jedes Blatt ist der Hash einer Transaction. Jeder innere Knoten ist der Hash seiner zwei verketteten Kinder. Die Wurzel des Baums (die Merkle Root) ist ein einzelner Hash, der sich auf jede Transaction im Block festlegt.

Wenn ein Block vier Transactions (A, B, C, D) enthält, sieht der Baum so aus: hash(A) und hash(B) werden zu hash(AB) kombiniert. hash(C) und hash(D) werden zu hash(CD) kombiniert. hash(AB) und hash(CD) werden zur Merkle Root kombiniert.

Wenn die Anzahl der Transactions ungerade ist, wird der letzte Hash dupliziert, um auf jeder Ebene eine gerade Anzahl zu erreichen.
