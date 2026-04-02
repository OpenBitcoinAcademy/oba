## Generación aleatoria de claves

Las primeras billeteras de Bitcoin generaban cada clave privada de forma independiente usando un generador de números aleatorios. Cada clave no tenía relación con las demás. La billetera las almacenaba todas en un archivo de base de datos.

Este enfoque tenía un problema serio de respaldo. Cada nueva clave requería un nuevo respaldo. Si un usuario generaba 100 claves y hacía respaldo después de la clave 50, las claves 51 a la 100 se perderían si el archivo de la billetera se corrompía.

Bitcoin Core originalmente pregeneraba un lote de 100 claves para reducir la frecuencia de respaldos necesarios. Pero el problema fundamental persistía: las claves aleatorias independientes son difíciles de gestionar.
