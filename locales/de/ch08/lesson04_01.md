## SIGHASH-Typen

Eine Signature bindet sich nicht standardmäßig an die gesamte Transaction. Das SIGHASH-Flag, das jeder Signature angehängt wird, legt fest, welche Teile der Transaction die Signature abdeckt.

**SIGHASH_ALL** (0x01) ist der Standard. Die Signature bindet sich an alle Inputs und alle Outputs. Jede Änderung an der Transaction macht die Signature ungültig.

**SIGHASH_NONE** (0x02) bindet sich an alle Inputs, aber keine Outputs. Jeder kann die Outputs nach dem Signieren ändern. Wird in seltenen kollaborativen Protokollen verwendet.

**SIGHASH_SINGLE** (0x03) bindet sich an alle Inputs, aber nur an den Output mit demselben Index wie der signierte Input. Andere Outputs können geändert werden.

Der ANYONECANPAY-Modifier lässt sich mit jedem dieser Typen kombinieren und erlaubt Signatures, die sich nur an einen einzelnen Input binden.
