## Pagos condicionales a traves de saltos

Los Contratos con Bloqueo de Hash y Tiempo (HTLCs) permiten pagos entre partes que no comparten un canal directo. Un HTLC combina dos condiciones: un hashlock (revelar un secreto para reclamar fondos) y un timelock (si el secreto no se revela a tiempo, los fondos regresan al remitente).

Carol genera un secreto aleatorio R y envia el hash H(R) a Alice en una factura de pago. Alice no conoce R. Crea un HTLC con Bob: "Paga 1.001 BTC si revelas la preimagen de H(R) dentro de 40 bloques." Bob crea un HTLC con Carol: "Paga 1.000 BTC si revelas la preimagen de H(R) dentro de 30 bloques."
