## Tres Camadas: SimplPedPop, EncPedPop, ChillDKG

A base e o SimplPedPop. Cada participante gera seu proprio polinomio aleatorio de grau t-1 e envia uma avaliacao secreta para cada outro participante, junto com um commitment dos coeficientes de seu polinomio. Cada participante soma as avaliacoes que recebeu para calcular seu share secreto final. A chave publica do grupo e derivada da soma dos commitments de todos os participantes para seus termos constantes. Nenhuma parte segura o segredo completo.

SimplPedPop assume um canal seguro entre cada par de participantes. EncPedPop adiciona isso fazendo cada participante gerar um par de chaves de criptografia efemero e criptografar suas avaliacoes secretas com a chave publica do destinatario. Agora o protocolo funciona por um canal de broadcast publico, porque bisbilhoteiros nao podem descriptografar os shares secretos.

ChillDKG envolve o EncPedPop com uma camada de gerenciamento de sessao. Introduz uma chave secreta de host que cada participante mantem de forma persistente, um dataset de recuperacao comum que e identico para todos os participantes e um protocolo para detectar e tratar mau comportamento. A chave secreta de host, combinada com os dados de recuperacao comuns, permite que um participante reconstrua seu share se perder seu dispositivo de assinatura.
