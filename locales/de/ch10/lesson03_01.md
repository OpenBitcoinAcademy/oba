## Wie Blocks sich verbreiten

Wenn ein Miner einen gültigen Block findet, sendet er den Block an seine Peers. Diese Peers validieren ihn und leiten ihn an ihre Peers weiter. Der Block breitet sich in wenigen Sekunden über das Netzwerk aus.

Geschwindigkeit zählt. Wenn zwei Miner fast gleichzeitig gültige Blocks finden, hat der Block, der mehr Nodes zuerst erreicht, einen Vorteil. Miner, die zuerst von Block A erfahren, bauen auf Block A auf. Miner, die zuerst von Block B erfahren, bauen auf Block B auf. Das Rennen entscheidet sich, wenn der nächste Block eine Chain verlängert und der andere Block veraltet (stale) wird.

Langsame Verbreitung schadet der Dezentralisierung. Wenn große Mining Pools mit schnellen Verbindungen zuerst von Blocks erfahren, haben sie einen Vorsprung beim nächsten Block. Kleinere Miner mit langsameren Verbindungen verschwenden öfter Arbeit an veralteten Blocks. Das Netzwerkprotokoll begegnet dieser Asymmetrie mit Compact Block Relay.
