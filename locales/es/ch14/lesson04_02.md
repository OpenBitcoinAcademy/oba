## Limites eliminados

El Script heredado y SegWit v0 imponen limites estrictos de recursos: tamano de script de 10,000 bytes, 201 opcodes que no son push, 100 elementos en la pila del testigo. Estos limites eran necesarios porque el modelo de costo no consideraba scripts individuales.

Tapscript elimina el limite de tamano de script y el limite de conteo de opcodes. En su lugar, usa un presupuesto de operaciones de firma por entrada: 50 sigops + el peso del testigo de la entrada. Testigos mas grandes pagan comisiones mas altas, y el presupuesto de sigops escala con la comision pagada. El incentivo economico reemplaza el limite fijo.

Esto permite scripts que eran imposibles bajo los limites anteriores. Un umbral con 100 participantes, una cascada compleja con timelocks o una verificacion de cadena de hashes ahora pueden expresarse en una sola hoja Tapscript.
