## La frase de paso opcional

BIP39 soporta una frase de paso opcional que se mezcla en la derivación de la semilla. Las mismas 12 palabras con frases de paso diferentes producen semillas diferentes y billeteras diferentes.

Esto provee negación plausible: un usuario puede crear una billetera señuelo con una frase de paso y una billetera real con otra. Bajo coerción, puede revelar la frase de paso del señuelo sin exponer los fondos principales.

El riesgo: olvidar la frase de paso significa perder acceso a la billetera. No existe mecanismo de recuperación. La frase de paso no se almacena en ningún lugar. Existe solo en la memoria del usuario o en un respaldo físico.
