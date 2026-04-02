## Tabaka Tatu: Sera, Miniscript, Script

Miniscript inafanya kazi katika usanifu wa tabaka tatu. Juu ni lugha ya sera inayosomeka na binadamu. Katikati ni Miniscript yenyewe. Chini ni Bitcoin Script.

Sera inaelezea unachotaka: "Alice NA Bob, AU Carol baada ya siku 90." Unaandika hii kama `or(and(pk(Alice),pk(Bob)),and(pk(Carol),older(12960)))`. Lugha ya sera ni kwa binadamu. Inaweka majina kwa funguo na kufuli za wakati bila kujali opcode.

Mkusanyaji unatafsiri sera kuwa usemi wa Miniscript: `or_d(and_v(v:pk(Alice),pk(Bob)),and_v(v:pk(Carol),older(12960)))`. Usemi wa Miniscript unabainisha kwa usahihi jinsi masharti yanavyoundwa, ikijumuisha aina zipi za vipande zinatumika katika kila nafasi. Kutoka hapo, usemi unaoana moja kwa moja na opcode za Bitcoin Script. Kila kipande cha Miniscript kina usimbaji mmoja na mmoja tu wa Script.
