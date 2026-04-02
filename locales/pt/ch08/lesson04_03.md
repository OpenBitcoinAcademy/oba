## O Algoritmo de Assinatura do SegWit

O SegWit introduziu um novo algoritmo de assinatura (BIP143) que corrige um problema antigo: em transacoes legacy, o valor sendo gasto nao e incluido nos dados assinados. Isso forcava assinantes a buscar a transacao anterior completa para verificar o valor, desacelerando hardware wallets.

O BIP143 inclui o valor de cada input nos dados que sao assinados. Uma hardware wallet pode verificar o valor total de input e a taxa sem baixar transacoes anteriores. O processo de assinatura e mais rapido e mais seguro.

Para segwit v1 (Taproot), o BIP341 define um algoritmo atualizado que inclui compromissos adicionais, suportando tanto key path quanto script path para gastos.
