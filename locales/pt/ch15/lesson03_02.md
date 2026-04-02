## Descriptors e Miniscript Juntos

Descriptors suportam incorporar expressoes Miniscript. O descriptor `wsh()` envolve uma expressao Miniscript em um output P2WSH. O descriptor `tr()` coloca uma expressao Miniscript dentro de uma folha de arvore de scripts Taproot.

Um descriptor como `wsh(and_v(v:pk(Alice),or_d(pk(Bob),older(12960))))` define uma politica de gasto completa como uma unica string. Qualquer carteira que entende o formato de descriptor pode importar essa string, derivar os enderecos corretos, identificar as condicoes de gasto e construir transacoes validas. O descriptor carrega a politica completa, nao uma dica parcial.

E isso que torna carteiras interoperaveis. Um assinante de hardware pode exibir as condicoes de gasto parseadas do descriptor. Uma carteira somente-leitura pode monitorar os enderecos resultantes. Um coordenador de assinatura pode construir um PSBT que satisfaz a politica. Cada ferramenta le a mesma string de descriptor e chega ao mesmo output de Script.

Descriptors incluem um checksum: oito caracteres anexados apos um simbolo `#`. O checksum detecta erros de transcricao. Uma carteira rejeita qualquer descriptor cujo checksum nao corresponda. Um unico caractere alterado na chave ou politica produz um checksum diferente, capturando erros antes que fundos sejam enviados para o endereco errado.
