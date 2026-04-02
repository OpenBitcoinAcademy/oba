## Lagrange-Interpolation

FROST baut auf Shamirs Secret Sharing auf. Während der Key-Generierung wird der Group Secret Key in n Shares aufgeteilt, sodass beliebige t Shares ihn rekonstruieren können. Das mathematische Werkzeug, das dies ermöglicht, ist die Lagrange-Interpolation.

Ein Polynom vom Grad t-1 ist durch t Punkte eindeutig bestimmt. Jeder Unterzeichner hält einen Punkt (seinen Secret Share) auf einem Polynom, dessen konstanter Term das Group Secret ist. Beim Signieren rekonstruieren die Unterzeichner nie das vollständige Polynom. Stattdessen multipliziert jeder Unterzeichner seinen Signature Share mit einem Lagrange-Koeffizienten, der von der Menge der teilnehmenden Unterzeichner abhängt.

Der Lagrange-Koeffizient für Unterzeichner $i$ in einer Signiermenge $S$ ist: $\lambda_i = \prod_{j \in S, j \neq i} \frac{j}{j - i}$. Diese Koeffizienten stellen sicher, dass die Summe der gewichteten Signature Shares eine gültige Schnorr-Signature für den Group Public Key ergibt, ohne dass jemals eine Partei den vollständigen Secret Key besitzt oder rekonstruiert.
