## Source-Based Routing

Der Absender, nicht das Netzwerk, wählt den Zahlungspfad. Jeder Lightning Node pflegt eine lokale Sicht der Netzwerktopologie (welche Nodes existieren, welche Channels sie verbinden, deren Kapazitäten und Gebührensätze). Der Absender berechnet eine Route aus diesen Informationen.

Das unterscheidet sich vom Internet-Routing, bei dem jeder Router den nächsten Hop eigenständig bestimmt. Bei Lightning hat der Absender die volle Kontrolle über den Pfad. Die Zwischenstationen folgen Weiterleitungsanweisungen, ohne die gesamte Route zu kennen.
