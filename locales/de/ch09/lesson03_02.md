## Child Pays for Parent (CPFP)

CPFP ist eine Technik zum Erhöhen der Fee, bei der der Empfänger eine neue Transaction erstellt, die den unbestätigten Output ausgibt. Die neue Transaction (das Child) zahlt eine hohe Fee-Rate. Miner, die die Fee des Child wollen, müssen auch den Parent aufnehmen.

CPFP erfordert keine Mitwirkung des ursprünglichen Absenders. Der Empfänger erstellt die Child-Transaction mit seinen eigenen Keys. Das macht CPFP nützlich, wenn der Absender nicht erreichbar ist.

Die kombinierte Fee-Rate von Parent und Child muss attraktiv genug sein, damit ein Miner beide aufnimmt. Miner bewerten Transaction-"Pakete" (Gruppen abhängiger Transactions), wenn sie auswählen, welche sie aufnehmen.
