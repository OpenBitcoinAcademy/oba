## O Diretorio de Dados

O Bitcoin Core armazena seus dados em um diretorio especifico da plataforma: `~/.bitcoin` no Linux, `~/Library/Application Support/Bitcoin` no macOS, `%APPDATA%\Bitcoin` no Windows.

O subdiretorio `blocks/` contem os arquivos de dados brutos dos blocos. O diretorio `chainstate/` guarda o banco de dados de UTXOs, um armazenamento LevelDB de todos os outputs nao gastos. O arquivo `mempool.dat` persiste o mempool entre reinicializacoes.

O arquivo de configuracao `bitcoin.conf` controla configuracoes de rede, limites de recursos, autenticacao RPC e flags de recursos. Configuracoes como `prune=550` habilitam o modo pruned, e `txindex=1` constroi um indice de todas as transacoes para buscas mais rapidas.
