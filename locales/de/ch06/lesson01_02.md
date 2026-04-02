## Inputs verweisen auf vorherige Outputs

Jeder Input zeigt auf einen bestimmten Output einer früheren Transaction. Er tut das mit zwei Informationen: der Transaction-ID (ein Hash) und dem Output-Index (welcher Output in dieser Transaction).

Wenn du Bitcoin ausgibst, beweist du, dass du den Key kontrollierst, der einen vorherigen Output entsperren kann. Bei Legacy-Transactions befindet sich der Beweis (Signature und Public Key) im Input Script. Bei modernen Segwit-Transactions befindet sich der Beweis in einer separaten Witness-Struktur, und das Input Script ist leer.

Sobald ein Output von einem Input referenziert wird, ist er ausgegeben. Er kann nicht erneut ausgegeben werden. So verhindert Bitcoin Double Spending ohne eine zentrale Autorität.
