## Zukünftige Upgrades über OP_SUCCESS

Tapscript reserviert eine Reihe von Opcodes namens OP_SUCCESSx. Enthält ein Script einen OP_SUCCESS-Opcode, ist das gesamte Script sofort erfolgreich, ohne weitere Validierung.

Das ist ein bewusster Upgrade-Mechanismus. Ein zukünftiger Soft Fork kann einen OP_SUCCESSx-Opcode neu definieren, um neue Validierung durchzuführen (wie OP_CHECKTEMPLATEVERIFY oder OP_VAULT). Alte Nodes sehen OP_SUCCESS und akzeptieren die Transaction. Neue Nodes führen den neu definierten Opcode aus und setzen die neuen Regeln durch.

Taproot unterstützt auch unbekannte Leaf-Versionen. Das aktuelle Tapscript verwendet Leaf-Version 0xC0. Ein zukünftiger Soft Fork kann Leaf-Version 0xC2 mit völlig anderer Script-Semantik definieren. Alte Nodes überspringen die Validierung unbekannter Leaf-Versionen. Neue Nodes setzen die neuen Regeln durch.

Beide Mechanismen erlauben es Taproot, sich weiterzuentwickeln, ohne ersetzt zu werden. Jede Script Leaf-Version und jeder OP_SUCCESS-Opcode ist ein reservierter Slot für zukünftige Funktionalität.
