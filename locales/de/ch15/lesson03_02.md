## Descriptors und Miniscript zusammen

Descriptors unterstützen das Einbetten von Miniscript-Ausdrücken. Der `wsh()`-Descriptor verpackt einen Miniscript-Ausdruck in einen P2WSH Output. Der `tr()`-Descriptor platziert einen Miniscript-Ausdruck in einem Taproot Script Tree Leaf.

Ein Descriptor wie `wsh(and_v(v:pk(Alice),or_d(pk(Bob),older(12960))))` definiert eine vollständige Ausgabepolicy als einzelnen String. Jedes Wallet, das das Descriptor-Format versteht, kann diesen String importieren, die korrekten Addresses ableiten, die Ausgabebedingungen erkennen und gültige Transactions konstruieren. Der Descriptor trägt die vollständige Policy, nicht einen Teilhinweis.

Das macht Wallets interoperabel. Ein Hardware Signer kann die aus dem Descriptor geparsten Ausgabebedingungen anzeigen. Ein Watch-Only-Wallet kann die resultierenden Addresses überwachen. Ein Signing Coordinator kann eine PSBT konstruieren, die die Policy erfüllt. Jedes Werkzeug liest denselben Descriptor-String und kommt zum selben Script Output.

Descriptors enthalten eine Prüfsumme: acht Zeichen, angehängt nach einem `#`-Symbol. Die Prüfsumme erkennt Übertragungsfehler. Ein Wallet lehnt jeden Descriptor ab, dessen Prüfsumme nicht übereinstimmt. Ein einzelnes geändertes Zeichen im Key oder der Policy erzeugt eine andere Prüfsumme und fängt Fehler ab, bevor Mittel an die falsche Address gesendet werden.
