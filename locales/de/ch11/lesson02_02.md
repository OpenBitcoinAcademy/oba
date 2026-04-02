## Die Kette der Hashes

Jeder Block-Header enthält den Hash des Headers des vorherigen Blocks. Block 500.000 enthält den Hash von Block 499.999, der den Hash von Block 499.998 enthält, und so weiter bis zum Genesis Block (Block 0).

Das Ändern von Daten in Block 499.999 würde seinen Hash ändern. Block 500.000 würde dann einen falschen Previous Block Hash enthalten und ungültig werden. Um einen historischen Block zu ändern, muss ein Angreifer den Proof of Work für diesen Block und jeden darauf folgenden Block wiederholen.

Diese kumulative Verknüpfung macht die Blockchain manipulationssicher. Je tiefer ein Block liegt, desto mehr Arbeit ist nötig, um ihn zu ändern. Jeder neue Block fügt eine weitere Schutzschicht für jeden Block darunter hinzu.
