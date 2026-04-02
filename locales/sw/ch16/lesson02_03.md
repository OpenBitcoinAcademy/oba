## Mchanganyaji, Mkamilishaji, Mtaji

Mchanganyaji anaunganisha PSBT nyingi zenye saini za sehemu kwa muamala uleule. Katika multisig ya 2-ya-3, msainiaji A anazalisha PSBT yenye saini yake na msainiaji B anazalisha nyingine. Mchanganyaji anachukua zote mbili, anaunganisha ingizo za PARTIAL_SIG kwa kila ingizo, na kuzalisha PSBT moja yenye saini zote zinazopatikana.

Mkamilishaji anabadilisha saini za sehemu kuwa scriptSig kamili na ushahidi kwa kila ingizo. Kwa ingizo ya P2PKH, anachukua PARTIAL_SIG moja na kujenga scriptSig. Kwa multisig ya P2WSH ya 2-ya-3, anachukua saini za sehemu, anazipanga kwa usahihi, na kuunda rundo la ushahidi pamoja na hati ya kukomboa. Baada ya kukamilisha, ramani za ingizo za PSBT zina sehemu za FINAL_SCRIPTSIG na FINAL_SCRIPTWITNESS. Data ya sehemu haihitajiki tena.

Mtaji anasoma PSBT iliyokamilishwa na kuzalisha muamala ghafi wa mtandao. Anachukua muamala ambao haujasainiwa kutoka ramani ya jumla, anaingiza FINAL_SCRIPTSIG na FINAL_SCRIPTWITNESS kutoka kila ingizo, na kusimba matokeo. Matokeo ni muamala wa kawaida wa Bitcoin ulio tayari kutangazwa. PSBT imetimiza madhumuni yake.
