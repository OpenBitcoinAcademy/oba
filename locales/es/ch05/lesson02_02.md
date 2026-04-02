## Generación determinista de claves

La generación determinista de claves resuelve el problema del respaldo. Una sola semilla aleatoria produce una secuencia de claves privadas mediante hashing repetido. Cada clave en la secuencia se puede recrear a partir de la semilla.

Respalda la semilla una vez, y puedes recuperar cada clave que la billetera ha generado o generará. La semilla es el secreto maestro. Si la pierdes, pierdes el acceso a todas las claves derivadas.

La forma más simple: empieza con una semilla, aplica hash para obtener la clave 1, aplica hash a la clave 1 para obtener la clave 2, y así sucesivamente. Esto produce una lista plana de claves. Pero las billeteras modernas usan un enfoque más sofisticado: un árbol.
