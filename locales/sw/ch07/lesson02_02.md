## Hati za Saini Nyingi

Hati ya saini nyingi (multisig) inahitaji saini M kutoka kwa funguo N za umma kuidhinisha kutumia. Multisig ya 2-ya-3 inahitaji saini mbili yoyote kati ya funguo tatu zilizoteuliwa.

Katika matokeo ya multisig wazi, hati ya kufunga ina: M <pubkey1> <pubkey2> ... <pubkeyN> N OP_CHECKMULTISIG. Hati ya kufungua inatoa: OP_0 <sig1> <sig2>. OP_0 ya mbele ni suluhisho la hitilafu ya kihistoria ya kuondoa-kwa-moja katika OP_CHECKMULTISIG.

Kwa vitendo, multisig inafungwa katika P2SH. Mtumaji anaona anwani ya kawaida ya P2SH. Maelezo ya multisig yamefichwa ndani ya hati ya kukomboa na yanafichuliwa tu wakati wa kutumia. Hii inashikilia matokeo kuwa madogo na ugumu uwe wa siri hadi wakati wa kutumia.
