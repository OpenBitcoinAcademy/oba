## Validacion del lado del cliente

En RGB, solo las partes de un contrato validan sus transiciones de estado. Un emisor de tokens y los titulares de tokens verifican el historial de transferencias. Ningun otro nodo en la red necesita procesar o almacenar estos datos.

Compara esto con sistemas donde cada nodo valida cada contrato (el modelo de Ethereum). La validacion del lado del cliente escala sin limite: agregar mas contratos no aumenta la carga sobre ningun nodo que no este involucrado. La privacidad es inherente: los datos del contrato son visibles solo para los participantes.

El costo: los participantes deben almacenar y verificar el historial completo de sus activos. Si el historial se pierde, el activo no puede verificarse. RGB usa cadenas de compromisos y pruebas para hacer esta verificacion eficiente.
