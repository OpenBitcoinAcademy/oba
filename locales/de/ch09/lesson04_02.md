## Transaction Pinning

Transaction Pinning ist ein Angriff, bei dem eine böswillige Partei eine Child-Transaction mit niedriger Fee erstellt, die es der ehrlichen Partei teuer oder unmöglich macht, die Fee des Parent zu erhöhen.

In einem Zwei-Parteien-Protokoll (wie einem Lightning-Network-Channel) könnte eine Partei einen großen Nachfahren mit niedriger Fee einer gemeinsamen Transaction senden. Der CPFP-Versuch der anderen Partei müsste auch für den großen Nachfahren des Angreifers zahlen, was die Fee-Erhöhung unerschwinglich teuer macht.

Anchor Outputs sind eine Gegenmaßnahme. Jede Partei erhält einen kleinen Output in der gemeinsamen Transaction, den sie mit CPFP ausgeben kann. Regeln begrenzen, wie viele Nachfahren jeder Anchor haben darf, und verhindern so den Pinning-Angriff.
