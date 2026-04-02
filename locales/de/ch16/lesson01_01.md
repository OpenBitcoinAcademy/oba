## Das Problem vor PSBTs

Eine Bitcoin Transaction benötigt Informationen aus mehreren Quellen, bevor sie gesendet werden kann. Das Wallet, das die Transaction erstellt, muss wissen, welche UTXOs ausgegeben werden. Der Signer muss die Private Keys besitzen. In vielen Setups sind das verschiedene Geräte oder verschiedene Software.

Vor BIP 174 gab es keinen Standard, um eine unsignierte Transaction zwischen diesen Beteiligten weiterzugeben. Jede Wallet-Software erfand ihr eigenes Format. Bitcoin Core serialisierte partielle Transactions anders als Electrum, das sie anders serialisierte als Hardware Wallets. Eine in einem Werkzeug erstellte Transaction konnte nicht von einem anderen signiert werden, ohne speziellen Verbindungscode.

Das machte Multisignature-Setups mühsam. Jeder Signer brauchte kompatible Software. Die Koordination zwischen Laptop, Hardware Wallet und Cold-Storage-Rechner erforderte manuelle Schritte und Formatkonvertierungen, die Fehler einführten.
