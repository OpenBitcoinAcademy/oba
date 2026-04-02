## Den Script-Baum aufbauen

Der Merkle Tree wird aus TapLeaf- und TapBranch-Hashes konstruiert.

Jedes Leaf: tagged_hash("TapLeaf", leaf_version || compact_size(script_length) || script). Die Leaf-Version für aktuelles Tapscript ist 0xC0.

Jeder Branch: tagged_hash("TapBranch", sorted(left_hash, right_hash)). Die beiden Kinder werden vor dem Hashing lexikographisch sortiert, sodass der Baum unabhängig von der Einfügereihenfolge eine einzige kanonische Form hat.

Der Baum kann bis zu 128 Ebenen tief sein, was bis zu 2^128 Leaves ermöglicht. Praktische Bäume sind weit kleiner. Scripts in geringer Tiefe (näher an der Root) haben kürzere Merkle Proofs und kosten weniger beim Ausgeben. Platziere den wahrscheinlichsten Rückfall auf Tiefe 1. Platziere selten genutzte Notfallpfade tiefer.
