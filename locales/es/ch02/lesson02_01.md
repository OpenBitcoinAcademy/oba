## Alice le paga a Bob

Alice quiere comprar un producto en la tienda en línea de Bob. La página de pago de Bob muestra una opción de pago con Bitcoin, con un monto y una dirección Bitcoin.

Alice abre su aplicación de billetera. Esta escanea la blockchain en busca de salidas de transacciones no gastadas (UTXOs) bloqueadas a sus direcciones. Su billetera encuentra un UTXO de 0.015 BTC, más que suficiente para cubrir el precio de Bob de 0.01 BTC.

La billetera de Alice construye una transacción con una entrada (referenciando su UTXO de 0.015 BTC) y dos salidas: 0.01 BTC a la dirección de Bob, y 0.0049 BTC de vuelta a Alice como cambio. Los 0.0001 BTC restantes son la comisión de transacción.
