## Weight Units

Segwit introduced weight units to replace the byte-based size metric. Each byte of non-witness data costs 4 weight units. Each byte of witness data (signatures) costs 1 weight unit.

This discount incentivizes the use of segwit transactions, which store signatures in the witness structure. A segwit transaction uses less weight than a legacy transaction with the same number of inputs and outputs.

Virtual bytes (vbytes) convert weight to a byte-equivalent: vbytes = weight / 4. A transaction with 600 weight units has 150 vbytes. Fee rates expressed in sat/vB account for the segwit discount automatically.
