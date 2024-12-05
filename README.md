# Hyperliquid user fill syncer

Sync your hyperliquid fills with local database to avoid losing your transaction history.

## Features

- Download fills at launch via REST API (max 2000 txs)
- Start syncing fill via websocket realtime
- All your fills will be stored into local database, currenly only sqlite3 is supported.

## Usage

Help

```
hypersync --help
```

Run syncer

```
hypersync --address $wallet_address --verbose
```