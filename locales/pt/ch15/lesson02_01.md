## Tres Camadas: Policy, Miniscript, Script

O Miniscript opera em uma arquitetura de tres camadas. No topo esta uma linguagem de politica legivel por humanos. No meio esta o Miniscript em si. Na base esta o Bitcoin Script.

Uma politica descreve o que voce quer: "Alice E Bob, OU Carol apos 90 dias." Voce escreve isso como `or(and(pk(Alice),pk(Bob)),and(pk(Carol),older(12960)))`. A linguagem de politica e para humanos. Ela nomeia chaves e timelocks sem se preocupar com opcodes.

Um compilador traduz a politica em uma expressao Miniscript: `or_d(and_v(v:pk(Alice),pk(Bob)),and_v(v:pk(Carol),older(12960)))`. A expressao Miniscript especifica exatamente como as condicoes sao compostas, incluindo quais tipos de fragmento sao usados em cada posicao. A partir dai, a expressao mapeia diretamente para opcodes de Bitcoin Script. Cada fragmento Miniscript tem uma e apenas uma codificacao em Script.
