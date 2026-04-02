## Multisignature Scripts

Ein Multisignature-Script (Multisig) erfordert M Signatures von N Public Keys, um eine Ausgabe zu autorisieren. Ein 2-von-3-Multisig erfordert beliebige 2 von 3 festgelegten Keys.

In einem blanken Multisig-Output enthält das Locking Script: M <pubkey1> <pubkey2> ... <pubkeyN> N OP_CHECKMULTISIG. Das Unlocking Script liefert: OP_0 <sig1> <sig2>. Das führende OP_0 ist ein Workaround für einen historischen Off-by-one-Bug in OP_CHECKMULTISIG.

In der Praxis wird Multisig in P2SH verpackt. Der Sender sieht eine Standard-P2SH-Address. Die Multisig-Details sind im Redeem Script verborgen und werden erst beim Ausgeben enthüllt. Das hält Outputs kompakt und die Komplexität privat bis zum Zeitpunkt der Ausgabe.
