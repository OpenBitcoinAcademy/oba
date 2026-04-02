## Campos Taproot e Diferencas de Versao

O BIP 371 adicionou campos especificos do Taproot ao formato PSBT. TAP_KEY_SIG (tipo de chave de input 0x13) armazena uma assinatura Schnorr para gasto por key path. TAP_SCRIPT_SIG (tipo de chave de input 0x14) armazena uma assinatura Schnorr para uma folha de script especifica, indexada tanto pela chave publica X-only quanto pelo hash da folha.

TAP_LEAF_SCRIPT (tipo de chave de input 0x15) fornece o script, sua versao de folha e o control block necessario para gasto por script path. TAP_BIP32_DERIVATION (tipo de chave de input 0x16) estende o campo de derivacao BIP 32 padrao com uma lista de hashes de folha para os quais a chave pode assinar. TAP_INTERNAL_KEY (tipo de chave de input 0x17) registra a chave publica interna antes do ajuste.

No lado do output, TAP_INTERNAL_KEY (tipo de chave de output 0x05) e TAP_BIP32_DERIVATION (tipo de chave de output 0x07) permitem que assinantes verifiquem que outputs de troco Taproot pertencem a mesma carteira. O assinante pode re-derivar a chave ajustada a partir da chave interna e confirmar que corresponde ao scriptPubKey do output.

O PSBTv2 (BIP 370) difere do v0 em estrutura, nao em conceito. No v0, a transacao nao assinada vive no mapa global como um blob serializado. No v2, inputs e outputs sao descritos por campos por mapa: PREVIOUS_TXID, OUTPUT_INDEX, SEQUENCE para inputs; AMOUNT e SCRIPT para outputs. Constructors podem adicionar inputs e outputs incrementalmente sem re-serializar toda a transacao a cada vez.
