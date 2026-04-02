## Limites Removidos

Script legacy e SegWit v0 impoem limites rigorosos de recursos: tamanho de script de 10.000 bytes, 201 opcodes nao-push, 100 elementos na pilha de witness. Esses limites eram necessarios porque o modelo de custo nao contabilizava scripts individuais.

Tapscript remove o limite de tamanho de script e o limite de contagem de opcodes. Em vez disso, usa um orcamento de operacoes de assinatura por input: 50 sigops + o peso de witness do input. Witness maiores pagam taxas mais altas, e o orcamento de sigops escala com a taxa paga. O incentivo economico substitui o limite fixo.

Scripts que eram impossiveis sob os limites antigos agora funcionam. Um threshold com 100 participantes, uma cascata complexa com timelocks ou uma verificacao de cadeia de hash podem ser expressos em uma unica folha Tapscript.
