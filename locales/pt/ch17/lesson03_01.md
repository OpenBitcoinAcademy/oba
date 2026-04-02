## O Problema do Dealer Confiavel

A forma mais simples de configurar shares de chave FROST e com um dealer confiavel. Uma parte gera a chave secreta do grupo, avalia o polinomio de Shamir em n pontos e distribui um share para cada participante. Isso funciona, mas cria um ponto unico de falha: o dealer conhece a chave secreta completa. Se o dealer for comprometido, todos os fundos do grupo estao em risco.

Um dealer comprometido tambem pode distribuir shares inconsistentes, dando a alguns participantes dados invalidos que farao a assinatura falhar de forma imprevisivel. Os participantes nao tem como verificar se seus shares estao corretos sem um protocolo que aplique consistencia.

ChillDKG, especificado no BIP 445, elimina o dealer confiavel. E um protocolo de geracao de chaves distribuida: o grupo produz conjuntamente os shares de chave sem que nenhuma parte segure ou veja a chave secreta completa. O protocolo se constroi em tres camadas, cada uma adicionando uma garantia especifica.
