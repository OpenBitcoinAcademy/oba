## Der Penalty Close

Wenn eine Partei eine widerrufene (alte) Commitment Transaction sendet, kann die andere Partei das Revocation Secret verwenden, um das gesamte Channel-Guthaben zu beanspruchen. Der Betrüger verliert alle Mittel im Channel.

Das ist der ökonomische Durchsetzungsmechanismus. Das Senden eines alten Zustands ist erkennbar (die andere Partei überwacht die Blockchain auf widerrufene Commitments) und bestrafbar (Totalverlust der Mittel). Die Strafe macht Lightning Channels vertrauenswürdig, ohne dass Vertrauen zwischen den Parteien nötig ist.

Die ehrliche Partei muss online sein (oder einen Watchtower-Dienst haben, der in ihrem Namen überwacht), um eine widerrufene Commitment-Übertragung innerhalb des Timelock-Fensters zu erkennen und darauf zu reagieren.
