## Usambazaji wa Vizuizi Vilivyofupishwa

Iliyofafanuliwa katika **BIP 152**, usambazaji wa vizuizi vilivyofupishwa unapunguza kipimo data kinachohitajika kusambaza kizuizi kipya. Wazo kuu: miamala mingi katika kizuizi kipya tayari iko katika mempool ya nodi inayopokea.

Badala ya kutuma kizuizi kamili, nodi inayotangaza inatuma ujumbe wa `cmpctblock`. Huu una kichwa cha kizuizi, orodha ya vitambulisho vifupi vya miamala, na muamala wa coinbase. Vitambulisho vifupi vya miamala ni hash zilizofupishwa za baiti 6 zinazomruhusu mpokea kulinganisha miamala aliyonayo tayari.

Nodi inayopokea inaunda upya kizuizi kutoka mempool yake kwa kutumia vitambulisho vifupi. Ikiwa miamala yoyote inakosekana, inaiomba kwa ujumbe wa `getblocktxn` na kuipokea kwa jibu la `blocktxn`.

BIP 152 inafafanua hali mbili. Katika **hali ya kipimo data cha chini**, nodi kwanza inatangaza kizuizi na ujumbe wa `inv`. Rika inaomba kizuizi kilichofupishwa ikiwa ina hamu. Katika **hali ya kipimo data cha juu**, nodi inatuma kizuizi kilichofupishwa mara moja bila kusubiri ombi. Wachimbaji na nodi zilizounganishwa vizuri kawaida hutumia hali ya kipimo data cha juu kupunguza ucheleweshaji.

Usambazaji wa vizuizi vilivyofupishwa unapunguza kipimo data cha kuenea kwa asilimia 90 au zaidi kwa kizuizi cha kawaida. Kuenea kwa haraka kunawapa wachimbaji wadogo nafasi ya haki, na kuimarisha umadaraka.
