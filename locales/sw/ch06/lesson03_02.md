## Baiti kwa Baiti

**Toleo** (baiti 4): 01000000 inamaanisha toleo 1. 02000000 inamaanisha toleo 2 (inawezesha vikwazo vya mfuatano vya BIP68). Usimbaji wa little-endian: baiti ndogo zaidi inakuja kwanza.

**Ingizo**: idadi ya varint ikifuatiwa na kila ingizo. Kila ingizo ina hash ya muamala wa awali (baiti 32, iliyogeuzwa), faharasa ya matokeo (baiti 4), urefu wa hati ya ingizo (varint), hati ya ingizo (inayobadilika), na nambari ya mfuatano (baiti 4).

**Matokeo**: idadi ya varint ikifuatiwa na kila matokeo. Kila matokeo yana thamani katika satoshi (baiti 8, little-endian) na urefu wa hati ya matokeo (varint) ikifuatiwa na hati ya matokeo.

**Ushahidi** (SegWit pekee): kwa kila ingizo, idadi ya vipengele vya ushahidi ikifuatiwa na urefu na data ya kila kipengele. Ingizo za zamani zina vipengele sifuri vya ushahidi.

**Locktime** (baiti 4): kawaida 00000000. Thamani isiyo sifuri inazuia muamala kuchimbwa lini.

Kitambulisho cha muamala (txid) ni hash mbili ya SHA-256 ya muamala uliosimbwa, data ya ushahidi ikiondolewa (kwa SegWit) au ikijumuishwa (kwa zamani).
