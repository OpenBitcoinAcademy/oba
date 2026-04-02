## The Coinbase Transaction

Every block begins with a special transaction called the coinbase transaction. It has no inputs from previous transactions. Instead, it creates new bitcoin as a reward for the miner who found the block.

The coinbase transaction has one input with a null reference (no previous output being spent). The miner fills this input's script field with arbitrary data, including the extra nonce used during mining. Satoshi Nakamoto famously embedded a newspaper headline in the genesis block's coinbase: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."

The coinbase transaction's outputs pay the miner. The total output value must not exceed the block subsidy plus the sum of all transaction fees in the block. Any value left unclaimed is destroyed.
