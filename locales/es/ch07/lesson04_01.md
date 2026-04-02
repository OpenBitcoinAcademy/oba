## Taproot (P2TR)

Taproot (BIP341) es el tipo de salida mas nuevo, activado en noviembre de 2021 como segwit version 1. Una salida Taproot se compromete con una sola clave publica. El gasto puede ocurrir de dos formas.

**Ruta de clave**: el propietario firma directamente con la clave comprometida. En la blockchain, esto se ve identico a un gasto de firma unica. Ningun script se revela. Nadie puede distinguir si la salida tenia condiciones de gasto adicionales.

**Ruta de script**: el propietario revela un script de un arbol Merkle de scripts comprometidos en la salida. Esto permite condiciones complejas (multisig, timelocks, hash locks) mientras mantiene el caso comun (ruta de clave) compacto y privado.

Las direcciones Taproot comienzan con "bc1p" en mainnet y usan codificacion bech32m.
