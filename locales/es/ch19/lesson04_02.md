## BitVM: verificacion de computacion

BitVM permite la verificacion de computacion Turing-completa en Bitcoin sin cambios de consenso. Un probador afirma que una funcion f(x) = y. Si la afirmacion es correcta, nada ocurre en cadena. Si es incorrecta, cualquier verificador puede enviar una prueba de fraude y el probador pierde un deposito.

BitVM2 (2024-2025) hizo la verificacion sin permisos: cualquiera puede desafiar una afirmacion falsa, y las disputas se resuelven en tres transacciones en cadena. El sistema implementa un verificador SNARK en Bitcoin Script, permitiendo la verificacion eficiente de computaciones complejas.

Las aplicaciones incluyen puentes sin confianza (mover bitcoin entre cadenas), rollups (agrupar transacciones fuera de la cadena con verificacion en cadena) y cualquier protocolo que se beneficie de "verificar, no computar" en Bitcoin.
