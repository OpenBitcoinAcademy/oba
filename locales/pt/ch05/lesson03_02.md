## Caminhos de Derivacao (BIP44)

O BIP44 define uma estrutura padrao para a arvore de chaves. O formato do caminho e: m / purpose' / coin_type' / account' / change / address_index.

**purpose** e 44 para enderecos padrao (ou 84 para SegWit, 86 para Taproot). **coin_type** e 0 para Bitcoin mainnet, 1 para testnet. **account** permite separar fundos em grupos logicos. **change** e 0 para enderecos de recebimento e 1 para enderecos de troco. **address_index** incrementa para cada novo endereco.

Um caminho tipico de endereco de recebimento Bitcoin: m/84'/0'/0'/0/0. Isso significa: proposito SegWit, Bitcoin mainnet, primeira conta, recebimento (nao troco), primeiro endereco.

Caminhos padronizados permitem que diferentes softwares de carteira reconstruam o mesmo conjunto de chaves a partir da mesma semente.
