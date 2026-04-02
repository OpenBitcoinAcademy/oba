What happens if a FROST signer reuses the same nonce in two different signing sessions?

- The final signature becomes invalid but no secret data is leaked
- An attacker can extract that signer's secret key share from the two signature shares
- The coordinator detects the reuse and aborts both sessions
