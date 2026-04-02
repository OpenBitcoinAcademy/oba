## Privacidade por Indistinguibilidade

O ganho de privacidade das assinaturas threshold vai alem de esconder a politica de gasto. Com OP_CHECKMULTISIG, o proprio template de script e uma impressao digital. Empresas de analise de cadeia categorizam enderecos por tipo de script, identificando carteiras multisig e inferindo arranjos de custodia.

FROST elimina essa impressao digital. Uma carteira de custodia 2-de-3, um tesouro corporativo 5-de-7 e uma carteira pessoal de chave unica produzem outputs on-chain identicos. Cada um gasta com uma chave publica de 32 bytes e uma assinatura de 64 bytes dentro de um key path Taproot.

Essa indistinguibilidade beneficia todos os usuarios Taproot. Quanto maior o conjunto de transacoes que parecem iguais, mais dificil e distinguir qualquer transacao individual. Usuarios de assinatura threshold se misturam na multidao de gastadores de chave unica, e gastadores de chave unica ganham negacao plausivel de que podem ser assinantes threshold.
