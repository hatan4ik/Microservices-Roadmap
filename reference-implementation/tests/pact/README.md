# Pact-style contract tests (training)

This folder shows the *shape* of consumer-driven contracts.

Recommended workflow:
1) consumer defines contract for provider interaction
2) provider verifies it in CI
3) publish contract to a broker (optional)

See `contract.json` for a simple example.

## How to use this in practice

- Treat the contract as a build artifact: version it and verify it automatically.
- Make breaking changes visible early:
  - consumers run a contract test in their CI to produce/update the contract
  - providers run verification in their CI against the latest contract(s)
- Keep contracts focused on behavior that matters (don’t snapshot every field “just because”).
