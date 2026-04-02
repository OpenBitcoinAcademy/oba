Que sucede si un firmante FROST reutiliza el mismo nonce en dos sesiones de firma diferentes?

- La firma final se vuelve invalida pero no se filtra ningun dato secreto
- Un atacante puede extraer la porcion de clave secreta de ese firmante a partir de las dos porciones de firma
- El coordinador detecta la reutilizacion y aborta ambas sesiones
