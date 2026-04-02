## Fedimint: Community-Verwahrung

Fedimint ist ein Protokoll für gemeinschaftlich betriebene Verwahrung mittels föderiertem Chaumian Ecash. Eine Gruppe von Guardians (vertrauenswürdige Community-Mitglieder) betreibt gemeinsam eine Mint per Threshold Multisig. Nutzer zahlen Bitcoin ein und erhalten im Gegenzug Ecash-Notes.

Die Ecash-Notes werden mit Blind Signatures erstellt. Beim Minting blendet der Nutzer eine zufällige Token-Kennung, bevor er sie an die Guardians sendet. Die Guardians signieren sie, ohne die Kennung zu sehen. Wenn der Nutzer das signierte Token entblendet, kann die Mint die Einzahlung nicht mit der Einlösung verknüpfen. Transactions zwischen Nutzern sind nicht verknüpfbar.

Nutzer vertrauen dem Threshold der Federation (eine Mehrheit der Guardians müsste sich verschwören, um zu stehlen). Das ist schwächer als Bitcoins vertrauensloses Modell, aber weit stärker, als einem einzelnen Verwahrer zu vertrauen. Fedimint zielt auf Communities mit bestehendem sozialem Vertrauen: Spargruppen, Kooperativen, lokale Organisationen.
