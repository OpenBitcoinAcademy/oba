## Aufgehobene Limits

Legacy Script und SegWit v0 setzen strikte Ressourcenlimits: 10.000 Bytes Script-Größe, 201 Non-Push-Opcodes, 100 Witness-Stack-Elemente. Diese Limits waren nötig, weil das Kostenmodell einzelne Scripts nicht berücksichtigte.

Tapscript hebt das Script-Größenlimit und das Opcode-Zähllimit auf. Stattdessen verwendet es ein Signature-Operations-Budget pro Input: 50 Sigops + das Witness-Gewicht des Inputs. Größere Witnesses verursachen höhere Gebühren, und das Sigops-Budget skaliert mit der gezahlten Gebühr. Der ökonomische Anreiz ersetzt das feste Limit.

Scripts, die unter den alten Limits unmöglich waren, funktionieren jetzt. Ein Schwellenwert mit 100 Teilnehmern, eine komplexe Timelock-Kaskade oder eine Hash-Chain-Verifikation lassen sich in einem einzigen Tapscript Leaf ausdrücken.
