## MuSig2 als Internal Key

Der Internal Key P kann ein MuSig2-Aggregat mehrerer Public Keys sein. Wenn Alice und Bob ihre Keys mittels MuSig2 zu P aggregieren, committet der getweakte Key Q auf beide. Bei Kooperation erzeugen sie eine einzige Schnorr Signature auf Q.

On-chain sieht diese 2-of-2 Multisig identisch aus wie eine Einzelsigner-Transaction. Der Output ist 34 Bytes. Die Witness ist 64 Bytes. Kein Beobachter kann feststellen, dass zwei Parteien beteiligt waren.

BitGo berichtete, dass ein MuSig2 Key Path Input 57,5 virtuelle Bytes kostet, verglichen mit 104,5 vbytes für einen nativen SegWit P2WSH Multisig Input. Die Einsparungen entstehen durch den Wegfall der Einzel-Public-Keys und Signatures, die Multisig Scripts on-chain erfordern.
