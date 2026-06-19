# Nebula Chain - Lokalna sieć Cosmos SDK

Nazwa: **Nebula** (NEB)
Denom: `unebula` (1 NEB = 1,000,000 unebula)
Chain ID: `nebula-1`
Adresy: `nebula1...`

## Struktura projektu

```
/root/nebula/
  app/              # Główna aplikacja Cosmos SDK
  cmd/nebulad/      # Binary węzła (nebulad)
  config.yml        # Konfiguracja Ignite CLI
  docs/             # Dokumentacja API
  proto/            # Pliki protobuf
  testutil/         # Test helpers
```

## Uruchomienie

```bash
# Inicjalizacja (jeśli potrzeba od nowa)
rm -rf ~/.nebula
nebulad init my-moniker --chain-id nebula-1

# Dodanie kont
nebulad keys add alice --keyring-backend test
nebulad keys add bob --keyring-backend test

# Dodanie do genesis
nebulad genesis add-genesis-account alice 100000000000unebula --keyring-backend test
nebulad genesis add-genesis-account bob 50000000000unebula --keyring-backend test

# Stworzenie walidatora
nebulad genesis gentx alice 50000000000unebula --keyring-backend test --chain-id nebula-1
nebulad genesis collect-gentxs

# Konfiguracja denom in genesis (stake -> unebula)
jq '.app_state.crisis.constant_fee.denom = "unebula" |
    .app_state.gov.params.min_deposit[0].denom = "unebula" |
    .app_state.mint.params.mint_denom = "unebula" |
    .app_state.staking.params.bond_denom = "unebula"' \
    ~/.nebula/config/genesis.json > tmp.json && mv tmp.json ~/.nebula/config/genesis.json

# Uruchomienie
nebulad start
```

## Przydatne komendy

```bash
# Status węzła
nebulad status

# Balans konta
nebulad query bank balances nebula1...

# Wysyłanie tokenów
nebulad tx bank send alice nebula1... 1000000unebula --keyring-backend test --chain-id nebula-1

# Staking (delegacja)
nebulad tx staking delegate nebulavaloper1... 1000000unebula --from alice --keyring-backend test --chain-id nebula-1

# Nagrody walidatora
nebulad query distribution rewards nebula1... --chain-id nebula-1

# Propozycja governance
nebulad tx gov submit-proposal --title "Test" --description "Test" --deposit 10000000unebula --from alice --chain-id nebula-1 --keyring-backend test
```

## Porty (domyślnie)
- RPC: 26657
- P2P: 26656
- gRPC: 9090
- REST API: 1317
