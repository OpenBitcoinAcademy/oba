## UTXOs auswählen

Das Erstellen einer Transaction beginnt mit der Auswahl der auszugebenden Outputs. Dein Wallet durchsucht die Blockchain nach UTXOs, die an deine Addresses gebunden sind. Das sind die Coins, die du kontrollierst.

Um 0,5 BTC zu senden, wählt dein Wallet einen oder mehrere UTXOs aus, die zusammen mindestens 0,5 BTC ergeben. Wenn der kleinste UTXO 0,8 BTC beträgt, wird dieser einzelne UTXO zum Input. Der Überschuss (abzüglich Gebühren) kommt als Wechselgeld zu dir zurück.

Wenn kein einzelner UTXO groß genug ist, kombiniert das Wallet mehrere UTXOs als separate Inputs in derselben Transaction. Jeder Input braucht seine eigene Signature.
