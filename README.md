# argentum_go
A simple technical analisys tool implementation

- Project
  - v1
    - Fetch Trades from Exchanges (Kraken/Binance)
      - (extra) store data into a DB or a Redis Cache
      - (extra) allow websocket connection to keep fetching/consuming new trades
    - Map trade data into a unifed trade model
    - use the Trade model to create Candlestick models
    - allow the user to specifiy the candle interval (strating from 1 minute)
  - v2
    - web app to render candlestick chart using chartJS
    - allow live updates via websocket or server sent event (SSE)
