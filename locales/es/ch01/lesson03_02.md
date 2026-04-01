## De clave privada a clave publica

A partir de una clave privada, puedes calcular una clave publica. Esto usa multiplicacion de curva eliptica, una operacion matematica facil de realizar en una direccion y computacionalmente inviable de revertir.

La clave privada es un numero. Multiplicalo por un punto especifico en una curva conocida (llamado punto generador G), y obtienes otro punto en la curva. Ese resultado es tu clave publica.

Cualquiera puede verificar que una clave publica corresponde a una clave privada comprobando la matematica. Pero dada solo la clave publica, nadie puede determinar la clave privada. Esta relacion unidireccional es la base de la seguridad de Bitcoin.
