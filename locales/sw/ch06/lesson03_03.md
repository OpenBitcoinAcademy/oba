## Varint na Endianness

Bitcoin inatumia kanuni mbili za usimbaji katika usimbaji wake wote.

**Varint** (nambari kamili za urefu unaobadilika) zinasimba idadi na urefu. Thamani chini ya 253 inafaa katika baiti moja. Thamani 253-65535 zinatumia kiambishi cha 0xFD ikifuatiwa na baiti 2. Thamani hadi takriban bilioni 4 zinatumia 0xFE ikifuatiwa na baiti 4. Thamani kubwa zaidi zinatumia 0xFF ikifuatiwa na baiti 8.

Mpangilio wa baiti wa **little-endian** unaweka baiti ndogo zaidi kwanza. Toleo 1 linasimbwa kama 01 00 00 00, si 00 00 00 01. Thamani za satoshi, urefu wa vizuizi, na sehemu nyingi za nambari zinatumia little-endian.

Hash za miamala (txid) zinaonyeshwa kwa mpangilio wa baiti uliogeuzwa kwa kanuni. Hash mbili ya SHA-256 inazalisha baiti kwa mpangilio mmoja, lakini vivinjari vya vizuizi na Bitcoin Core vinavyoonyesha kwa kugeuzwa. Hii ni kanuni ya kuonyesha, si sheria ya usimbaji.
