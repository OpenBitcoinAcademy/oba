## Package Relay

Historicamente, o Bitcoin Core avaliava cada transacao individualmente para admissao no mempool. Uma transacao mae de taxa baixa seria rejeitada mesmo que uma transacao filha de taxa alta dependesse dela.

O package relay muda isso. Os nos avaliam grupos de transacoes relacionadas juntos, aceitando uma mae de taxa baixa se sua filha trouxer a taxa combinada acima do limite.

Isso melhora a confiabilidade do CPFP. Sem package relay, a mae poderia nao se propagar para os mineradores, tornando a filha inutil. Com package relay, mae e filha viajam juntas pela rede.
