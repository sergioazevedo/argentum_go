# argentum_go
A simple technical analisys tool implementation

how to run
```
make clean run-console
```

![alt text](chart.png "Candlesticks chart")

- Project
  - v1 ✅
    - Fetch Trades from Exchanges (Kraken/Binance) ✅
    - Map trade data into a unifed trade model ✅
    - use the Trade model to create Candlestick models ✅
    - allow the user to specifiy the candle interval (strating from 1 minute) ✅
    - render candlestick charts into html files from console ✅

Maybe for later
  - v2
    - web app to render candlestick chart using chartJS
    - allow live updates via websocket or server sent event (SSE)
  - v3
    - (extra) store data into a DB or a Redis Cache (v3 maybe)
    - (extra) allow websocket connection to keep fetching/consuming new trades (v4 maybe)

