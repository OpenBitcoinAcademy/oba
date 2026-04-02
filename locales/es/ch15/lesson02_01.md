## Tres capas: politica, Miniscript, Script

Miniscript opera en una arquitectura de tres capas. En la parte superior hay un lenguaje de politicas legible por humanos. En el medio esta Miniscript. En la parte inferior esta Bitcoin Script.

Una politica describe lo que quieres: "Alice Y Bob, O Carol despues de 90 dias." Se escribe como `or(and(pk(Alice),pk(Bob)),and(pk(Carol),older(12960)))`. El lenguaje de politicas es para humanos. Nombra claves y timelocks sin preocuparse por los opcodes.

Un compilador traduce la politica a una expresion Miniscript: `or_d(and_v(v:pk(Alice),pk(Bob)),and_v(v:pk(Carol),older(12960)))`. La expresion Miniscript especifica exactamente como se componen las condiciones, incluyendo que tipos de fragmentos se usan en cada posicion. A partir de ahi, la expresion se mapea directamente a opcodes de Bitcoin Script. Cada fragmento Miniscript tiene una y solo una codificacion en Script.
