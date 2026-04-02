## BitVM: Berechnungsverifikation

BitVM ermöglicht Turing-vollständige Berechnungsverifikation auf Bitcoin ohne Konsensänderungen. Ein Prover behauptet, dass eine Funktion f(x) = y ergibt. Wenn die Behauptung korrekt ist, passiert on-chain nichts. Wenn sie falsch ist, kann jeder Verifier einen Fraud Proof einreichen und der Prover verliert eine Kaution.

BitVM2 (2024-2025) machte die Verifikation permissionless: Jeder kann eine falsche Behauptung anfechten, und Streitigkeiten werden in drei On-Chain-Transactions gelöst. Das System implementiert einen SNARK-Verifier in Bitcoin Script und ermöglicht so die effiziente Verifikation komplexer Berechnungen.

Anwendungen umfassen vertrauenslose Bridges (Bitcoin zwischen Chains bewegen), Rollups (Transactions off-chain bündeln mit On-Chain-Verifikation) und jedes Protokoll, das von "verifizieren statt berechnen" auf Bitcoin profitiert.
