## PSBT ni Nini

Partially Signed Bitcoin Transaction (PSBT) ni muundo wa binary uliofafanuliwa na BIP 174. Inafunga muamala ambao haujasainiwa pamoja na metadata ambayo kila mshiriki anahitaji kufanya kazi yake. Waundaji wanaambatanisha data ya UTXO. Wasasishaji wanaongeza njia za uundaji. Wasainiaji wanaongeza saini. Muundo unabeba kila kitu kinachohitajika ili kila jukumu liweze kufanya kazi kwa uhuru.

Binary inaanza na uchawi wa baiti tano: `0x70736274FF`. Baiti nne za kwanza zinaandika "psbt" kwa ASCII. Baiti ya tano, `0xFF`, inaweka alama ya kitenganishi. Zana yoyote inayopokea blob inaweza kukagua baiti hizi tano kuthibitisha ni PSBT kabla ya kuchambua zaidi.

Baada ya uchawi inakuja mfuatano wa ramani za funguo-thamani. Kila ingizo la ramani lina aina ya funguo, mzigo wa funguo, na mzigo wa thamani. Baiti ya sifuri (0x00) inakamilisha kila ramani. Ramani ya kwanza ni ramani ya jumla. Ramani zinazofuata zinapishana kati ya ramani za kila ingizo na kila matokeo, moja kwa kila ingizo na matokeo katika muamala.
