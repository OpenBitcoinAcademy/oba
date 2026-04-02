## Privatsphäre durch selektive Offenlegung

Ein Taproot Output mit 8 Script Leaves legt beim Ausgeben über den Script Path nur ein Leaf offen. Die anderen 7 Leaves erscheinen als Hashes im Merkle Proof. Ein Beobachter erfährt eine Ausgabebedingung, kann aber nicht feststellen, was die anderen Bedingungen waren.

Vergleiche das mit P2WSH Multisig: beim Ausgeben werden das vollständige Script, alle Public Keys und welche Keys signiert haben offengelegt. Jede beteiligte Partei ist sichtbar.

Ein 2-of-3 Multisig mit Taproot: der Key Path enthält das MuSig2-Aggregat der zwei häufigsten Signer. Zwei Script Leaves enthalten die Rückfallpaare (A+C und B+C). Im Normalfall wird der Key Path genutzt und nichts über die Multisig offengelegt. Bei einem Rückfall wird nur ein alternatives Paar offengelegt. Das andere Paar bleibt ein Hash.
