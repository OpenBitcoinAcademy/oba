## Roles, Not People

The PSBT workflow defines seven roles: Creator, Updater, Signer, Combiner, Finalizer, Extractor, and (in PSBTv2) Constructor. Each role performs one step. A single person or program can fill multiple roles, but the separation matters because each role has different trust requirements and different access to information.

The Creator builds the initial PSBT. In v0 (BIP 174), the Creator produces the unsigned transaction and wraps it in a PSBT with empty input and output maps. In v2 (BIP 370), the Creator sets global fields like the transaction version and locktime, but does not yet include inputs or outputs. That task falls to the Constructor.

The Constructor role exists only in PSBTv2. It adds inputs and outputs to the PSBT. Multiple Constructors can collaborate: one adds the inputs it controls, another adds its inputs, and each adds the outputs they need. This enables interactive transaction construction where no single party knows the full transaction shape until all Constructors have contributed.
