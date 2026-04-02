## Rutas de derivación (BIP44)

BIP44 define una estructura estándar para el árbol de claves. El formato de ruta es: m / purpose' / coin_type' / account' / change / address_index.

**purpose** es 44 para direcciones estándar (o 84 para SegWit, 86 para Taproot). **coin_type** es 0 para Bitcoin mainnet, 1 para testnet. **account** permite separar fondos en grupos lógicos. **change** es 0 para direcciones de recepción y 1 para direcciones de cambio. **address_index** incrementa con cada nueva dirección.

Una ruta típica de dirección de recepción de Bitcoin: m/84'/0'/0'/0/0. Esto significa: propósito SegWit, Bitcoin mainnet, primera cuenta, recepción (no cambio), primera dirección.

Las rutas estandarizadas permiten que diferentes software de billetera reconstruyan el mismo conjunto de claves a partir de la misma semilla.
