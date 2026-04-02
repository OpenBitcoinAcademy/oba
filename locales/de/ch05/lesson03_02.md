## Ableitungspfade (BIP44)

BIP44 definiert eine Standardstruktur für den Key-Baum. Das Pfadformat lautet: m / purpose' / coin_type' / account' / change / address_index.

**purpose** ist 44 für Standard-Addresses (oder 84 für SegWit, 86 für Taproot). **coin_type** ist 0 für Bitcoin Mainnet, 1 für Testnet. **account** erlaubt die Trennung von Mitteln in logische Gruppen. **change** ist 0 für Empfangs-Addresses und 1 für Wechselgeld-Addresses. **address_index** zählt für jede neue Address hoch.

Ein typischer Bitcoin-Empfangs-Address-Pfad: m/84'/0'/0'/0/0. Das bedeutet: SegWit-Zweck, Bitcoin Mainnet, erstes Konto, Empfang (kein Wechselgeld), erste Address.

Standardisierte Pfade ermöglichen verschiedenen Wallet-Programmen, denselben Satz von Keys aus demselben Seed zu rekonstruieren.
