## Die 51%-Attacke und Double Spending

Ein Angreifer, der mehr als die Hälfte der Hash-Leistung des Netzwerks kontrolliert, kann die ehrliche Chain überholen. Der Angreifer mined Blocks im Geheimen und baut einen privaten Zweig auf, der länger ist als der öffentliche. Nachdem der Händler eine Zahlung mit mehreren Bestätigungen akzeptiert hat, veröffentlicht der Angreifer den privaten Zweig. Das Netzwerk reorganisiert sich zur längeren Chain, und die Version des Angreifers ersetzt das Original. Die Zahlung verschwindet.

Das ist eine Double-Spend-Attacke. Der Angreifer gibt dieselben Coins zweimal aus: einmal an den Händler auf der öffentlichen Chain und einmal zurück an sich selbst auf der privaten Chain. Die Kosten sind enorm. Der Angreifer muss für die Dauer der Attacke die Mehrheit der Hash-Leistung aufrechterhalten und verliert die Block-Belohnungen auf der aufgegebenen öffentlichen Chain.

Mit weniger als 50 % der Hash-Leistung sinkt die Erfolgswahrscheinlichkeit exponentiell mit jeder Bestätigung, die der Verteidiger abwartet. Sechs Bestätigungen machen eine Attacke mit weniger als der Mehrheit statistisch vernachlässigbar.
