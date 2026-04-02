## Phrases de recuperation BIP39

BIP39 encode une graine de portefeuille sous forme d'une sequence de 12 ou 24 mots anglais, appelee phrase de recuperation (ou phrase de graine). Chaque mot provient d'une liste standardisee de 2 048 mots.

L'encodage par mots sert deux objectifs. Les mots sont plus faciles a ecrire, relire et transcrire que de l'hexadecimal brut. Et le dernier mot inclut une somme de controle qui detecte les erreurs de transcription.

Une phrase de 12 mots encode 128 bits d'entropie. Une phrase de 24 mots encode 256 bits. Les deux sont suffisamment fortes pour les besoins de securite actuels. La phrase est convertie en graine de 512 bits en utilisant PBKDF2 avec 2 048 cycles de HMAC-SHA512.
