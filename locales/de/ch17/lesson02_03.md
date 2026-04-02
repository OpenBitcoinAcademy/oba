## Von Shares zur Standard-Signature

Der Coordinator sammelt t Signature Shares und summiert sie. Das Ergebnis ist eine einzelne Schnorr-Signature: ein 64-Byte-(R, s)-Paar, wobei R der aggregierte Nonce-Punkt und s die Summe der Signature Shares ist. Diese Signature ist unter dem aggregierten Group Public Key gültig.

Ein Bitcoin Full Node, der diese Transaction verifiziert, führt die Standard-BIP-340-Schnorr-Verifikationsgleichung aus. Er prüft eine Signature gegen einen Public Key. Der Node hat keine Möglichkeit zu wissen, ob die Signature von drei Unterzeichnern, fünf Unterzeichnern oder einem Unterzeichner erzeugt wurde. Die Verifikation ist identisch.

Darum ist FROST so wirkungsvoll für Bitcoin: Es erfordert keine Konsensänderungen. Taproot und BIP 340 akzeptieren die Signatures, die FROST erzeugt, bereits. Die Threshold-Signing-Komplexität liegt vollständig in der Software der Unterzeichner. Die Blockchain und alle Verifizierer bleiben ahnungslos.
