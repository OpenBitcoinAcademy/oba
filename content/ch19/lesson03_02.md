## Cashu: Single-Mint Ecash

Cashu is a Chaumian ecash protocol using a single mint operator instead of a federation. Users deposit bitcoin (typically via Lightning) and receive blind-signed ecash tokens. The tokens are bearer instruments: whoever holds a valid token can redeem it.

The mint cannot link minting to redemption (blind signatures ensure this). The mint cannot see who transfers tokens to whom. Transfers between users are instant and free (no blockchain interaction).

The trade-off: users trust a single mint operator. If the mint is dishonest or goes offline, funds may be lost. Cashu is designed for small-scale, application-specific deployments (a coffee shop, a streaming service, a community) where the trust relationship is manageable.
