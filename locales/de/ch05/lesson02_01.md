## Zufällige Key-Erzeugung

Die frühesten Bitcoin-Wallets erzeugten jeden Private Key unabhängig mit einem Zufallszahlengenerator. Jeder Key hatte keinen Bezug zu den anderen. Das Wallet speicherte alle in einer Datenbankdatei.

Dieser Ansatz hatte ein gravierendes Backup-Problem. Jeder neue Key erforderte ein neues Backup. Wenn ein Nutzer 100 Keys erzeugte und nach Key 50 ein Backup machte, gingen die Keys 51 bis 100 verloren, falls die Wallet-Datei beschädigt wurde.

Bitcoin Core erzeugte ursprünglich einen Pool von 100 Keys im Voraus, um die Häufigkeit nötiger Backups zu reduzieren. Das Problem blieb bestehen: Unabhängige zufällige Keys sind schwer zu verwalten.
