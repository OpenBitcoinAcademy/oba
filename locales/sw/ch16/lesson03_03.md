## Sehemu za Taproot na Tofauti za Toleo

BIP 371 iliongeza sehemu maalum za Taproot kwenye muundo wa PSBT. TAP_KEY_SIG (aina ya funguo ya ingizo 0x13) inahifadhi saini ya Schnorr kwa kutumia kwa njia ya funguo. TAP_SCRIPT_SIG (aina ya funguo ya ingizo 0x14) inahifadhi saini ya Schnorr kwa jani maalum la hati, iliyofungukiwa na funguo ya umma ya X pekee na hash ya jani.

TAP_LEAF_SCRIPT (aina ya funguo ya ingizo 0x15) inatoa hati, toleo lake la jani, na kizuizi cha udhibiti kinachohitajika kwa kutumia kwa njia ya hati. TAP_BIP32_DERIVATION (aina ya funguo ya ingizo 0x16) inaongeza sehemu ya kawaida ya uundaji wa BIP 32 na orodha ya hash za majani ambayo funguo inaweza kusaini. TAP_INTERNAL_KEY (aina ya funguo ya ingizo 0x17) inarekodi funguo ya umma ya ndani kabla ya marekebisho.

Kwa upande wa matokeo, TAP_INTERNAL_KEY (aina ya funguo ya matokeo 0x05) na TAP_BIP32_DERIVATION (aina ya funguo ya matokeo 0x07) zinawapa wasainiaji uwezo wa kuthibitisha kwamba matokeo ya chenji ya Taproot ni ya pochi ileile. Msainiaji anaweza kutoa upya funguo iliyorekebishwa kutoka funguo ya ndani na kuthibitisha inalingana na scriptPubKey ya matokeo.

PSBTv2 (BIP 370) inatofautiana na v0 katika muundo, si dhana. Katika v0, muamala ambao haujasainiwa uko katika ramani ya jumla kama blob moja iliyosimbwa. Katika v2, ingizo na matokeo zinaelezwa na sehemu za kila ramani: PREVIOUS_TXID, OUTPUT_INDEX, SEQUENCE kwa ingizo; AMOUNT na SCRIPT kwa matokeo. Wajenzi wanaweza kuongeza ingizo na matokeo taratibu bila kusimba upya muamala mzima kila wakati.
