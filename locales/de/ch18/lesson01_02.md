## Commitment Transactions

Jede Partei hält ihre eigene Version der aktuellen Commitment Transaction. Diese Transactions teilen die Channel-Mittel gemäß dem aktuellen Guthaben auf. Wenn Alice 0,7 BTC hat und Bob 0,3 BTC, spiegeln beide Commitment Transactions diese Aufteilung wider.

Commitment Transactions sind asymmetrisch. Alices Version zahlt Bob sofort aus, sperrt aber Alices Mittel hinter einem Timelock. Bobs Version macht das Gegenteil. Diese Asymmetrie ermöglicht den Bestrafungsmechanismus: Wenn Alice eine alte Commitment Transaction sendet (um mehr zu beanspruchen, als ihr zusteht), kann Bob ein Revocation Secret verwenden, um alle Mittel zu beanspruchen, bevor Alices Timelock abläuft.

Bei jeder Guthabenaktualisierung tauschen beide Parteien Revocation Secrets für die alten Commitments aus, was deren Broadcast unsicher macht.
