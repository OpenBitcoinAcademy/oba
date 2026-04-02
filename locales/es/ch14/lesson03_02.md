## Construir el arbol de scripts

El arbol de Merkle se construye a partir de hashes TapLeaf y TapBranch.

Cada hoja: tagged_hash("TapLeaf", leaf_version || compact_size(script_length) || script). La version de hoja para Tapscript actual es 0xC0.

Cada rama: tagged_hash("TapBranch", sorted(left_hash, right_hash)). Los dos hijos se ordenan lexicograficamente antes del hashing, asegurando que el arbol tenga una sola forma canonica sin importar el orden de insercion.

El arbol puede tener hasta 128 niveles de profundidad, permitiendo hasta 2^128 hojas. Los arboles practicos son mucho mas pequenos. Los scripts ubicados a poca profundidad (mas cerca de la raiz) tienen pruebas de Merkle mas cortas y cuestan menos para gastar. Coloca el respaldo mas probable a profundidad 1. Coloca las rutas de emergencia poco usadas a mayor profundidad.
