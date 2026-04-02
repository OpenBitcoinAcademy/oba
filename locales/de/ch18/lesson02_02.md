## Preimage-Propagation

Carol kennt R. Sie offenbart es Bob und beansprucht 1,000 BTC. Bob kennt nun R. Er offenbart es Alice und beansprucht 1,001 BTC. Das Preimage propagiert rückwärts entlang des Zahlungspfades.

Alice hat 1,001 BTC bezahlt. Carol hat 1,000 BTC erhalten. Bob hat 0,001 BTC als Routing-Gebühr verdient. Die Zahlung wurde in Sekunden abgewickelt, ohne die Blockchain zu berühren.

Alle HTLCs in der Kette verwenden denselben Hash H(R). Entweder propagiert das Preimage den gesamten Weg zurück (die Zahlung gelingt vollständig) oder die Timelocks laufen ab (die Zahlung scheitert vollständig, alle Mittel werden zurückgegeben). Kein Zwischenzustand ist möglich. Die Zahlung ist atomar.
