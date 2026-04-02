## Selecionando UTXOs

Construir uma transacao comeca escolhendo quais outputs gastar. Sua carteira varre o blockchain em busca de UTXOs vinculados aos seus enderecos. Essas sao as moedas que voce controla.

Para enviar 0,5 BTC, sua carteira seleciona um ou mais UTXOs que somem pelo menos 0,5 BTC. Se o menor UTXO e 0,8 BTC, esse unico UTXO se torna o input. O excesso (menos taxas) retorna para voce como troco.

Se nenhum UTXO individual for grande o suficiente, a carteira combina multiplos UTXOs como inputs separados na mesma transacao. Cada input requer sua propria assinatura.
