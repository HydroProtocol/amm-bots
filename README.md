# Amm-bots

## How to run?

### Run from source

```
BOT_TYPE=CONST_PRODUCT \
BOT_PRIVATE_KEY=0x............... \
BOT_BASE_TOKEN=HOT \
BOT_QUOTE_TOKEN=DAI \
BOT_BASE_URL=http://127.0.0.1:3001 \
BOT_MIN_PRICE=0.8 \
BOT_MAX_PRICE=1.2 \
BOT_PRICE_GAP=0.02 \
BOT_EXPAND_INVENTORY=2 \
BOT_WEB3_URL=http://127.0.0.1:8545 \
go run main.go
```

### Run from docker image

```
docker run --rm -it \
  -e BOT_TYPE=CONST_PRODUCT \
  -e BOT_PRIVATE_KEY=0x............... \
  -e BOT_BASE_TOKEN=HOT \
  -e BOT_QUOTE_TOKEN=DAI \
  -e BOT_BASE_URL=http://127.0.0.1:3001 \
  -e BOT_MIN_PRICE=0.8 \
  -e BOT_MAX_PRICE=1.2 \
  -e BOT_PRICE_GAP=0.02 \
  -e BOT_EXPAND_INVENTORY=2 \
  -e BOT_WEB3_URL=http://127.0.0.1:8545 \
  hydroprotocolio/amm-bots
```

## Environment Variables

### Basic Variables
 - `BOT_TYPE` Type of bot
 - `BOT_BASE_URL` Hydro relayer api base url
 - `BOT_WEB3_URL` Ethereum json rpc url 
 - `BOT_BASE_TOKEN` Symbol of base token *(e.g HOT)*
 - `BOT_QUOTE_TOKEN` Symbol of quote token *(e.g WETH)*
 
### Secret Variables

 - `BOT_PRIVATE_KEY` Ethereum address for place order

## Algorithm Variables

### Const Product AMM
[Mechanism of const product market making](https://medium.com/scalar-capital/uniswap-a-unique-exchange-f4ef44f807bf)

This bot discretizes the continuous price curve and create a limit number of orders. The order price is limited between `BOT_MAX_PRICE` and `BOT_MIN_PRICE`. The price difference between adjacent orders is `BOT_PRICE_GAP`.

Const product algorithm has a disadvantage of low inventory utilization. For example, it only use 5% of your inventory when the price increases 10%. `BOT_EXPAND_INVENTORY` can help you add depth near the current price.

 - `BOT_MAX_PRICE` Max order price
 - `BOT_MIN_PRICE` Min order price
 - `BOT_PRICE_GAP` Price difference rate between adjacent orders. For example, ask price increases by 2% and bid price decreases by 2% if `BOT_PRICE_GAP=0.02`.
 - `BOT_EXPAND_INVENTORY` Multiply your order size linearly. For example, all order size will be tripled if `BOT_EXPAND_INVENTORY=3`.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details

