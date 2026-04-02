## El cierre por penalizacion

Si una parte transmite una transaccion de compromiso revocada (antigua), la otra parte puede usar el secreto de revocacion para reclamar el saldo completo del canal. El tramposo pierde todos los fondos en el canal.

Este es el mecanismo de aplicacion economica. Transmitir un estado antiguo es detectable (la otra parte monitorea la blockchain buscando compromisos revocados) y sancionable (perdida total de fondos). La penalizacion hace que los canales Lightning sean confiables sin requerir confianza entre las partes.

La parte honesta debe estar en linea (o tener un servicio de atalaya monitoreando en su nombre) para detectar y responder a la transmision de un compromiso revocado dentro de la ventana de timelock.
