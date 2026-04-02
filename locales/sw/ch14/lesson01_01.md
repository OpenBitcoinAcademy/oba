## Funguo ya Ndani na Marekebisho

Matokeo ya Taproot yanafunga fedha kwa funguo ya umma iliyorekebishwa Q. Funguo hii inatokana na ingizo mbili: funguo ya umma ya ndani P na mzizi wa hiari wa Merkle wa mti wa hati.

Thamani ya marekebisho t inakokotolewa kama: t = tagged_hash("TapTweak", P || merkle_root). Funguo iliyorekebishwa ni: Q = P + t * G, ambapo G ni nukta ya jenereta kwenye secp256k1.

Ikiwa hakuna mti wa hati, marekebisho yanatumia P pekee: t = tagged_hash("TapTweak", P). Matokeo bado yanajifunga na funguo ya ndani, lakini hakuna hati zilizopachikwa.

Kwenye mnyororo, scriptPubKey ni: OP_1 ikifuatiwa na kuratibu X ya baiti 32 ya Q. Hii ni baiti 34, ukubwa sawa na matokeo ya P2WSH. Hakuna mwangalizi anayeweza kujua ikiwa Q ina mti wa hati uliopachikwa au la.
