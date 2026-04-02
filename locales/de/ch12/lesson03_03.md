## Forks und Reorganisationen

Zwei Miner finden manchmal fast gleichzeitig gültige Blocks. In diesem Fall sehen Teile des Netzwerks den einen Block zuerst und andere Teile den anderen. Die Chain gabelt sich vorübergehend in zwei konkurrierende Zweige.

Das ist ein natürlicher Fork, kein Angriff. Jede Node arbeitet an dem Zweig weiter, den sie zuerst erhalten hat. Der Gleichstand löst sich, wenn der nächste Block gefunden wird. Wenn ein Miner einen Zweig verlängert, hat dieser Zweig nun mehr kumulativen Proof of Work. Nodes auf dem kürzeren Zweig wechseln zum längeren. Dieser Wechsel heißt Reorganisation. Die Transactions im aufgegebenen Block kehren in den Mempool zurück, um in einem zukünftigen Block aufgenommen zu werden.

Die meisten natürlichen Forks lösen sich innerhalb eines Blocks. Tiefere Reorganisationen sind selten und werden mit jeder weiteren Bestätigung exponentiell unwahrscheinlicher.
