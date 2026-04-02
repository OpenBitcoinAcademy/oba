## Ark: UTXOs Off-Chain

Ark e um protocolo Layer 2 que habilita transacoes Bitcoin baratas sem exigir canais. Um Ark Service Provider (ASP) coordena rodadas onde usuarios trocam Virtual UTXOs (VTXOs): transacoes Bitcoin validas mantidas off-chain.

Cada VTXO e uma folha em uma arvore de transacoes enraizada em um unico UTXO compartilhado on-chain. Usuarios sempre podem transmitir seu VTXO para reivindicar fundos on-chain (saida unilateral). Durante operacao normal, o ASP agrupa milhares de transferencias em uma unica liquidacao on-chain.

Ark nao requer abrir canais ou bloquear liquidez antecipadamente. Usuarios recebem VTXOs cedendo antigos em rodadas periodicas. Trocas atomicas via connector outputs garantem que nenhum fundo e perdido durante a troca.
