## Sehemu za Kila Ingizo

Kila ramani ya ingizo inaelezea ingizo moja ya muamala. Sehemu muhimu ambazo msainiaji anahitaji ni data ya UTXO, njia ya uundaji wa funguo ya kusaini, na hati zozote zinazohitajika kuunda ushahidi.

WITNESS_UTXO (aina ya funguo 0x01) inahifadhi matokeo yanayotumika: kiasi katika satoshi na scriptPubKey. Kwa ingizo za SegWit, hii inatosha kwa kusaini kwa sababu algorithm ya sighash inajifunga na kiasi moja kwa moja. NON_WITNESS_UTXO (aina ya funguo 0x00) inahifadhi muamala mzima wa awali, muhimu kwa ingizo za zamani ambapo msainiaji lazima athibitishe kiasi kwa kutafuta matokeo katika muamala kamili.

BIP32_DERIVATION (aina ya funguo 0x06) inaoanisha funguo ya umma na njia yake ya uundaji ya BIP 32 na alama ya kidole ya funguo kuu. Msainiaji analinganisha alama ya kidole na funguo yake kuu, anatoa funguo binafsi kwenye njia iliyotolewa, na kusaini. PARTIAL_SIG (aina ya funguo 0x02) inahifadhi saini iliyofungukiwa na funguo ya umma iliyoizalisha. SIGHASH_TYPE (aina ya funguo 0x03) inabainisha bendera ipi ya sighash msainiaji atumie.

Kwa ingizo za P2SH, REDEEM_SCRIPT (aina ya funguo 0x04) inabeba hati ya kukomboa. Kwa ingizo za P2WSH, WITNESS_SCRIPT (aina ya funguo 0x05) inabeba hati ya ushahidi. Msainiaji anahitaji hizi kukokotoa sighash sahihi.
