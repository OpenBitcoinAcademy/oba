## Politicas de Gasto com Decaimento Temporal

Uma politica de recuperacao com timelock comeca com uma condicao de gasto primaria e adiciona condicoes alternativas que ativam apos um atraso. A chave primaria controla os fundos imediatamente. Se a chave primaria for perdida ou seu detentor ficar indisponivel, um caminho de gasto alternativo desbloqueia apos um numero especificado de blocos.

Em Miniscript, isso e expresso como: `or(pk(primary),and(pk(recovery),older(26280)))`. A chave primaria pode gastar a qualquer momento. A chave de recuperacao pode gastar somente apos aproximadamente seis meses (26.280 blocos a 10 minutos por bloco). A politica compila para um unico Script com dois caminhos de execucao.

Esse padrao resolve um problema real. Uma carteira de chave unica nao tem caminho de recuperacao. Se a chave for perdida, os fundos se foram. Um multisig padrao requer multiplos assinantes para cada transacao, adicionando atrito ao uso diario. Uma alternativa com timelock da ao detentor primario controle total durante a operacao normal enquanto garante que uma parte confiavel pode recuperar fundos apos um atraso.
