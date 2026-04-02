## Saraka ya Data

Bitcoin Core inahifadhi data yake katika saraka maalum ya jukwaa: `~/.bitcoin` kwenye Linux, `~/Library/Application Support/Bitcoin` kwenye macOS, `%APPDATA%\Bitcoin` kwenye Windows.

Saraka ndogo ya `blocks/` ina faili za data ghafi za vizuizi. Saraka ya `chainstate/` inashikilia hifadhidata ya UTXO, hifadhi ya LevelDB ya matokeo yote ambayo hayajatumika. Faili ya `mempool.dat` inahifadhi mempool kati ya kuanza upya.

Faili ya usanidi `bitcoin.conf` inasimamia mipangilio ya mtandao, vikomo vya rasilimali, uthibitishaji wa RPC, na bendera za vipengele. Mipangilio kama `prune=550` inawezesha hali ya kupunguzwa, na `txindex=1` inajenga faharasa ya miamala yote kwa utafutaji wa haraka zaidi.
