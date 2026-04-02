## Zwei-Runden-Signierung

Das FROST-Signieren geschieht in zwei Runden zwischen den teilnehmenden Unterzeichnern. Ein Coordinator leitet Nachrichten weiter, erfährt aber nie genug, um eine Signature zu fälschen oder geheime Shares zu extrahieren. Der Coordinator ist nicht vertrauenswürdig.

In Runde eins erzeugt jeder Unterzeichner ein Paar zufälliger Nonces und sendet die zugehörigen Nonce-Commitments an den Coordinator. Der Coordinator sammelt die Commitments von allen t teilnehmenden Unterzeichnern und verteilt sie an die Gruppe. Diese Commitments binden jeden Unterzeichner an seine Nonces, bevor jemand die Werte der anderen Unterzeichner sieht. So wird verhindert, dass ein bösartiger Unterzeichner seine Nonces adaptiv wählt, um das Ergebnis zu manipulieren.

In Runde zwei berechnet jeder Unterzeichner die aggregierte Nonce aus den gesammelten Commitments, konstruiert den Schnorr-Challenge-Hash und erzeugt einen Signature Share unter Verwendung seines Secret Key Share und seiner Nonce. Jeder Unterzeichner sendet seinen Signature Share an den Coordinator.
