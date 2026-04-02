## Propagacion de la preimagen

Carol conoce R. Lo revela a Bob y reclama 1.000 BTC. Bob ahora conoce R. Lo revela a Alice y reclama 1.001 BTC. La preimagen se propaga hacia atras a lo largo de la ruta de pago.

Alice pago 1.001 BTC. Carol recibio 1.000 BTC. Bob gano 0.001 BTC como comision de enrutamiento. El pago se liquido en segundos sin tocar la blockchain.

Todos los HTLCs en la cadena usan el mismo hash H(R). O la preimagen se propaga completamente hacia atras (el pago tiene exito en su totalidad) o los timelocks expiran (el pago falla en su totalidad, todos los fondos son devueltos). No hay estado intermedio posible. El pago es atomico.
