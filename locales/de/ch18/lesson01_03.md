## Der Revocation-Mechanismus

Wenn Alice und Bob ihr Channel-Guthaben aktualisieren, entwerten sie das vorherige Commitment durch Austausch von Revocation Secrets. Jede Partei hält nun ein Secret, mit dem sie die andere für das Senden des alten Zustands bestrafen kann.

Wenn Bob eine alte Commitment Transaction sendet, in der er 0,5 BTC hatte (das aktuelle Guthaben gibt ihm aber nur 0,3 BTC), kann Alice Bobs Revocation Secret verwenden, um das gesamte Channel-Guthaben zu beanspruchen. Bob verliert alles.

Diese Strafe macht Betrug wirtschaftlich irrational. Die Kosten eines Betrugsversuchs (Verlust aller Mittel) übersteigen den möglichen Gewinn (Beanspruchung eines etwas höheren alten Guthabens) bei weitem. Ehrliches Verhalten ist die dominante Strategie.
