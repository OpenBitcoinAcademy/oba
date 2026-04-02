## Forks and Reorganizations

Two miners sometimes find valid blocks at nearly the same time. When this happens, parts of the network see one block first, and other parts see the other. The chain temporarily forks into two competing branches.

This is a natural fork, not an attack. Each node works on whichever branch it received first. The tie breaks when the next block is found. If a miner extends one branch, that branch now has more cumulative proof of work. Nodes on the shorter branch switch to the longer one. This switch is called a reorganization. The transactions in the abandoned block return to the mempool for inclusion in a future block.

Most natural forks resolve within one block. Deeper reorganizations are rare and become exponentially less likely with each additional confirmation.
