## Nos Completos e Nos Podados

Um no completo que mantem todos os blocos desde o bloco genesis e chamado de **no de arquivo**. Ele pode servir dados historicos de blocos para qualquer par que os solicite. Novos nos entrando na rede dependem de nos de arquivo para baixar o blockchain completo.

Executar um no de arquivo requer centenas de gigabytes de espaco em disco. Em 2024, o blockchain excede 600 GB. Nem todo operador de no completo pode arcar com esse armazenamento.

Um **no podado** valida cada bloco, mas descarta dados antigos de blocos apos processa-los. Ele mantem apenas o conjunto de UTXOs e os blocos mais recentes. Um no podado aplica todas as regras de consenso. Nao depende de confianca. A troca: nao pode servir blocos historicos para outros pares.

Tanto nos de arquivo quanto nos podados sao nos completos. Ambos validam tudo. A diferenca e se armazenam blocos antigos para outros baixarem.
