## Das Trusted-Dealer-Problem

Der einfachste Weg, FROST Key Shares einzurichten, ist mit einem Trusted Dealer. Eine Partei erzeugt den Group Secret Key, wertet das Shamir-Polynom an n Punkten aus und verteilt einen Share an jeden Teilnehmer. Das funktioniert, schafft aber einen Single Point of Failure: Der Dealer kennt den vollständigen Secret Key. Wenn der Dealer kompromittiert wird, sind die gesamten Mittel der Gruppe gefährdet.

Ein kompromittierter Dealer kann auch inkonsistente Shares verteilen und einigen Teilnehmern ungültige Daten geben, die dazu führen, dass das Signieren unvorhersehbar fehlschlägt. Teilnehmer haben ohne ein Protokoll, das Konsistenz erzwingt, keine Möglichkeit zu überprüfen, ob ihre Shares korrekt sind.

ChillDKG, spezifiziert in BIP 445, beseitigt den Trusted Dealer. Es ist ein Distributed-Key-Generation-Protokoll: Die Gruppe erzeugt die Key Shares gemeinsam, ohne dass eine einzelne Partei jemals den vollständigen Secret Key hält oder sieht. Das Protokoll besteht aus drei Schichten, die jeweils eine bestimmte Garantie hinzufügen.
