## Propagacao de Preimagem

Carol conhece R. Ela o revela para Bob e reivindica 1,000 BTC. Bob agora conhece R. Ele o revela para Alice e reivindica 1,001 BTC. A preimagem se propaga para tras ao longo do caminho de pagamento.

Alice pagou 1,001 BTC. Carol recebeu 1,000 BTC. Bob ganhou 0,001 BTC como taxa de roteamento. O pagamento foi liquidado em segundos sem tocar o blockchain.

Todos os HTLCs na cadeia usam o mesmo hash H(R). Ou a preimagem se propaga por todo o caminho de volta (o pagamento tem sucesso inteiramente) ou os timelocks expiram (o pagamento falha inteiramente, todos os fundos devolvidos). Nenhum estado intermediario e possivel. O pagamento e atomico.
