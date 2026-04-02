## Compact Block Filters

O **BIP 157** e o **BIP 158** definem uma abordagem melhor chamada Compact Block Filters. Em vez de o cliente enviar seu filtro para o servidor, o servidor constroi um filtro para cada bloco e o envia ao cliente.

Cada filtro e uma representacao compacta de todos os enderecos e scripts em um bloco. O cliente baixa o filtro e verifica localmente se algum de seus enderecos aparece. Se uma correspondencia for encontrada, o cliente baixa o bloco completo para extrair as transacoes relevantes.

A vantagem de privacidade e que o servidor nunca descobre quais enderecos o cliente procura. O servidor envia o mesmo filtro para cada cliente. O cliente faz toda a correspondencia localmente.

O custo de largura de banda e maior que Bloom filters porque o cliente baixa um filtro para cada bloco. Mas os filtros sao pequenos (cerca de 20 KB por bloco em media) e podem ser verificados contra um compromisso de hash na cadeia de cabecalhos de bloco. O cliente nao precisa confiar no servidor para fornecer filtros precisos.

Executar o Bitcoin sobre **Tor** adiciona outra camada de privacidade. O Tor esconde o endereco IP do cliente dos nos aos quais ele se conecta. Combinado com Compact Block Filters, um cliente leve pode verificar seu saldo sem revelar sua identidade ou seus enderecos a nenhum par.
