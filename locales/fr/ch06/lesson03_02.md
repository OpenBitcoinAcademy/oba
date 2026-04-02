## Octet par octet

**Version** (4 octets) : 01000000 signifie version 1. 02000000 signifie version 2 (active les contraintes de sequence BIP68). Encodage little-endian : l'octet le moins significatif vient en premier.

**Entrees** : un varint compteur suivi de chaque entree. Chaque entree contient le hachage de la transaction precedente (32 octets, inverse), l'index de sortie (4 octets), la longueur du script d'entree (varint), le script d'entree (variable) et le numero de sequence (4 octets).

**Sorties** : un varint compteur suivi de chaque sortie. Chaque sortie contient la valeur en satoshis (8 octets, little-endian) et la longueur du script de sortie (varint) suivie du script de sortie.

**Temoin** (segwit uniquement) : pour chaque entree, un compteur d'elements temoins suivi de la longueur et des donnees de chaque element. Les entrees legacy ont zero elements temoins.

**Locktime** (4 octets) : generalement 00000000. Une valeur non nulle restreint le moment ou la transaction peut etre minee.

L'identifiant de transaction (txid) est le hachage double-SHA-256 de la transaction serialisee, avec les donnees temoin exclues (pour segwit) ou incluses (pour legacy).
