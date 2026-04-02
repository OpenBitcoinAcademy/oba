## Pagamentos Condicionais entre Saltos

Hash Time-Locked Contracts (HTLCs) permitem pagamentos entre partes que nao compartilham um canal direto. Um HTLC combina duas condicoes: um hashlock (revele um segredo para reivindicar fundos) e um timelock (se o segredo nao for revelado a tempo, os fundos retornam ao remetente).

Carol gera um segredo aleatorio R e envia o hash H(R) para Alice em uma fatura de pagamento. Alice nao conhece R. Ela cria um HTLC com Bob: "Pague 1,001 BTC se voce revelar a preimagem de H(R) dentro de 40 blocos." Bob cria um HTLC com Carol: "Pague 1,000 BTC se voce revelar a preimagem de H(R) dentro de 30 blocos."
