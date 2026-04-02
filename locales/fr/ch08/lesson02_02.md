## Serialisation DER

Les signatures ECDSA dans Bitcoin sont encodees en utilisant DER (Distinguished Encoding Rules), un format binaire standard. L'encodage DER enveloppe les valeurs r et s avec des octets de type et de longueur.

Une signature ECDSA typique encodee en DER fait 71 a 73 octets de long. Elle est suivie d'un drapeau SIGHASH d'un octet qui indique quelles parties de la transaction la signature couvre.

Les transactions segwit v0 utilisent toujours ECDSA mais exigent un encodage DER strict (BIP66). Cela a elimine une source de malleabilite de transaction ou la meme signature valide pouvait etre encodee de plusieurs manieres, produisant des identifiants de transaction differents.
