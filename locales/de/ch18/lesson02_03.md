## Timelock-Dekremente

Jeder Hop in der Kette hat einen kürzeren Timelock als der vorherige Hop. Alice gibt Bob 40 Blocks. Bob gibt Carol 30 Blocks. Dieses Dekrement (das CLTV-Delta) stellt sicher, dass jeder weiterleitende Node Zeit hat, Mittel von seinem Downstream-Hop zu beanspruchen, bevor sein Upstream-HTLC abläuft.

Wenn Carol R nicht innerhalb von 30 Blocks offenlegt, läuft Bobs HTLC ab und die Mittel gehen an Bob zurück. Bob hat noch 10 Blocks Zeit, um mit Alice abzurechnen. Wenn Bob R nicht innerhalb von 40 Blocks gegenüber Alice offenlegt, gehen Alices Mittel an sie zurück.

Das Dekrement verhindert einen Timing-Angriff, bei dem ein Downstream-Node die Offenlegung des Preimage verzögert, bis der Upstream-HTLC abläuft, und so den weiterleitenden Node in die Falle lockt.
