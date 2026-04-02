## Taproot (P2TR)

Taproot (BIP341) e o tipo de output mais recente, ativado em novembro de 2021 como segwit versao 1. Um output Taproot faz commit a uma unica chave publica. O gasto pode acontecer de duas formas.

**Key path**: o proprietario assina diretamente com a chave comprometida. No blockchain, isso parece identico a um gasto de assinatura unica. Nenhum script e revelado. Ninguem pode dizer se o output tinha condicoes de gasto adicionais.

**Script path**: o proprietario revela um script de uma arvore Merkle de scripts comprometidos no output. Condicoes complexas sao possiveis (multisig, timelocks, hash locks) enquanto mantem o caso comum (key path) compacto e privado.

Enderecos Taproot comecam com "bc1p" na mainnet e usam codificacao bech32m.
