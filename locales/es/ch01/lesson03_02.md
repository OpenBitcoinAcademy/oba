## De clave privada a clave publica

A partir de una clave privada, puedes calcular una clave publica. Esto usa multiplicacion de curva eliptica, una operacion matematica facil de realizar en una direccion y computacionalmente inviable de revertir.

La clave privada es un numero. Multiplicalo por un punto especifico en una curva conocida (llamado punto generador G), y obtienes otro punto en la curva. Ese resultado es tu clave publica.

Dada solo la clave publica, nadie puede determinar la clave privada. Cualquiera puede verificar una firma digital contra una clave publica, confirmando que el firmante controla la clave privada correspondiente, sin que la clave privada sea revelada. Esta relacion unidireccional es la base de la seguridad de Bitcoin.
