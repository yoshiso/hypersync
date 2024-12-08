# Hyperliquid user fill syncer

Sync your hyperliquid fills with local database to avoid losing your transaction history.

## Features

- Download fills at launch via REST API (max 2000 txs)
- Start syncing fill via websocket realtime
- All your fills will be stored into local database, currenly only sqlite3 is supported.
- By default, db file is created on ./db.sqlite3
- Also support fundings, genesisSpot, rewardsClaim, vault deposit/withdraw/create, withdraw, spot/internal transfers also required to calculate PnLs.

## Misc

Build portable binary, will be built on bin/hypersync

```
make build
```

## Install

```
make build
sudo mv bin/hypersync /usr/local/bin
sudo chmod +755 /usr/local/bin/hypersync
```

## Usage

Help

```
hypersync --help
```

Run syncer

```
hypersync --address $wallet_address --verbose
```

Run syncer with output path

```
hypersync --address $wallet_address --out path/to/database.sqlite3
```