## Bifurcaciones y reorganizaciones

Dos mineros a veces encuentran bloques validos casi al mismo tiempo. Cuando esto sucede, partes de la red ven un bloque primero, y otras partes ven el otro. La cadena se bifurca temporalmente en dos ramas competidoras.

Esta es una bifurcacion natural, no un ataque. Cada nodo trabaja en la rama que recibio primero. El empate se rompe cuando se encuentra el siguiente bloque. Si un minero extiende una rama, esa rama ahora tiene mas proof of work acumulado. Los nodos en la rama mas corta cambian a la mas larga. Este cambio se llama reorganizacion. Las transacciones en el bloque abandonado regresan al mempool para su inclusion en un bloque futuro.

La mayoria de las bifurcaciones naturales se resuelven en un bloque. Las reorganizaciones profundas son raras y se vuelven exponencialmente menos probables con cada confirmacion adicional.
