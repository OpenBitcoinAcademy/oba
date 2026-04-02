## Rollen, nicht Personen

Der PSBT-Workflow definiert sieben Rollen: Creator, Updater, Signer, Combiner, Finalizer, Extractor und (in PSBTv2) Constructor. Jede Rolle führt einen Schritt aus. Eine einzelne Person oder ein Programm kann mehrere Rollen ausfüllen, aber die Trennung ist relevant, weil jede Rolle unterschiedliche Vertrauensanforderungen und unterschiedlichen Zugang zu Informationen hat.

Der Creator baut die initiale PSBT. In v0 (BIP 174) erzeugt der Creator die unsignierte Transaction und verpackt sie in eine PSBT mit leeren Input- und Output-Maps. In v2 (BIP 370) setzt der Creator globale Felder wie Transaction-Version und Locktime, enthält aber noch keine Inputs oder Outputs. Diese Aufgabe fällt dem Constructor zu.

Die Constructor-Rolle existiert nur in PSBTv2. Sie fügt der PSBT Inputs und Outputs hinzu. Mehrere Constructors können zusammenarbeiten: einer fügt die Inputs hinzu, die er kontrolliert, ein anderer seine Inputs, und jeder fügt die Outputs hinzu, die er braucht. Das ermöglicht interaktive Transaction-Konstruktion, bei der keine einzelne Partei die vollständige Transaction-Struktur kennt, bis alle Constructors beigetragen haben.
