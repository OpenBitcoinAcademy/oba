## Der Taproot-Safety-Tweak

ChillDKG erzeugt einen Group Public Key, aber dieser Key kann nicht direkt als Taproot-Output-Key verwendet werden. BIP 341 verlangt, dass der Output Key Q eine getweakte Version eines internen Key P ist: Q = P + tagged_hash("TapTweak", P) * G für einen Key-Path-Only-Output.

BIP 445 spezifiziert, dass ChillDKG diesen Tweak als Teil der Key-Generierung anwendet. Das Protokoll berechnet den getweakten Group Public Key und passt die Key Shares an, sodass die Gruppe für den getweakten Key signieren kann, ohne zusätzliche Schritte beim Signieren. Der Share jedes Teilnehmers wird modifiziert, um den Tweak zu berücksichtigen, und der von ChillDKG zurückgegebene Group Public Key ist der finale Taproot-Output-Key.

Diese Integration ist relevant, weil der Tweak konsistent angewendet werden muss. Wenn Teilnehmer sich über den getweakten Key nicht einig sind, schlägt das Signieren fehl. Indem der Tweak in das DKG-Protokoll selbst eingebettet wird, garantiert ChillDKG, dass alle Teilnehmer denselben Output Key ableiten und Shares halten, die damit konsistent sind.
