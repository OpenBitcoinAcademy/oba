## Construindo a Arvore de Scripts

A arvore Merkle e construida a partir de hashes TapLeaf e TapBranch.

Cada folha: tagged_hash("TapLeaf", leaf_version || compact_size(script_length) || script). A versao de folha para o Tapscript atual e 0xC0.

Cada ramo: tagged_hash("TapBranch", sorted(left_hash, right_hash)). Os dois filhos sao ordenados lexicograficamente antes do hashing, garantindo que a arvore tenha uma unica forma canonica independente da ordem de insercao.

A arvore pode ter ate 128 niveis de profundidade, permitindo ate 2^128 folhas. Arvores praticas sao muito menores. Scripts colocados em profundidades rasas (mais perto da raiz) tem provas Merkle mais curtas e custam menos para gastar. Coloque a alternativa mais provavel na profundidade 1. Coloque caminhos de emergencia raramente usados mais fundo.
