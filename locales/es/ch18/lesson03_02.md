## Enrutamiento cebolla

Lightning envuelve las instrucciones de pago en capas de cifrado, como una cebolla. Cada nodo de reenvio descifra una capa, que revela solo el siguiente salto y el monto a reenviar. El nodo no aprende el remitente, el destinatario final ni la longitud total de la ruta.

Este modelo de privacidad (basado en SPHINX) significa que Bob, al reenviar un pago de Alice a Carol, solo sabe que Alice le envio algo y que debe reenviarlo a Carol. No sabe si Alice es el remitente original u otro nodo de reenvio. No sabe si Carol es la destinataria final.

La cebolla tiene un tamano fijo independientemente del numero de saltos, previniendo el analisis de longitud de ruta.
