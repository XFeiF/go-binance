[link](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api_CN.md)

|    end_point                | query_type |  service  | function |
|-----------------------------|------------|------------|----------|
| /api/v3/account             | GET        | account          | 获取账户信息 |
| /api/v3/depth               | GET        |  depth           | 获取深度数据 |
| /api/v3/exchangeInfo        | GET        | exchange_info    | 获取交易所信息 |
| /api/v3/klines              | GET        | kline            | 获取K线数据 |  
| /api/v3/order               |         | order      | 获取平均价格 |
| /api/v3/order/test          |        | order            | 测试下单接口 |
| /api/v3/order/oco           |        | order            | OCO订单 |
| /api/v3/openOrderList       | GET        | order            |  |
| /api/v3/openOrders          | GET        | order            | 下单，查询，撤销订单 |
| /api/v3/order               | GET        | order            | 下单，查询，撤销订单 |
| /api/v3/allOrders           | GET        | order            | 下单，查询，撤销订单 |
| /api/v3/order               | DELETE     | order            | 下单，查询，撤销订单 |
| /api/v3/orderList           | DELETE     | order            | 下单，查询，撤销订单 |
| /api/v3/openOrders          | DELETE     | order            | 下单，查询，撤销订单 |
| /api/v3/rateLimit/order     | GET        | rate_limit       | 查询未成交的订单数量 |  
| /api/v3/ping                | GET        | server             | 检测服务器是否可用 |  
| /api/v3/time                | GET        | server             | 获取服务器时间 |  
| /api/v3/ticker/bookTicker   | GET        | ticker            | |
| /api/v3/ticker/price        | GET        | ticker           | |
| /api/v3/ticker/24hr         | GET        | ticker           | |
| /api/v3/avgPrice            | GET        | ticker           | |
| /api/v3/ticker              | GET        | ticker           | |
| /api/v3/myTrades            | GET        | trade            | 获取某交易对的成交历史 |  
| /api/v3/historicalTrades    | GET        | trade            | |  
| /api/v3/aggTrades           | GET        | trade            | 获取聚合交易数据 |  
| /api/v1/trades              | GET        | trade            | 获取某个交易对的最新成交数据 |  
| /api/v3/ticker/tradingDay   | GET        | tradingDayTicker | 获取交易日行情 |  
| /api/v3/uiKlines            | GET        | uiKlines         | 获取K线数据 |  
| /api/v3/userDataStream      | POST,PUT,DELETE   | user_stream      | 建立用户数据流 |  






- account_service
- asset_detail_service
- asset_dividend_service
- bnb_burn_service
- c2c_service
- client
- convert_trade
- deposit_service
- depth_service  
- dust_log_service
- exchange_info_service
- fiat_service
- futures_algo_service
- futures_service
- interest_history_service
- internal_universal_transfer_service
- kline_service
    - "/api/v3/klines"  K线数据
- liquidity_pool_service
- margin_order_service
- margin_service
- order_service
    - "/api/v3/order" 订单（下单，查询，撤销）
    - "/api/v3/order/test" 测试订单接口
    - "/api/v3/openOrders" （GET当前账户所有挂单）（DELETE撤销单一交易对的所有挂单）
    - "/api/v3/openOrderList" （GET当前挂单列表）
    - "/api/v3/allOrders" （GET所有订单，包括历史）（USER_DATA）
    - "/api/v3/order/oco" OCO订单（TRADE）
    - "/api/v3/orderList" 订单列表
- pay_service
- rate_limit_service
- rebate  
- request
- savings_service
- server_service
- staking_service
- subaccount_service
- ticker_service
    - "/api/v3/ticker"  # 滚动窗口价格变动统计
    - "/api/v3/ticker/bookTicker"  最优挂单接口
    - "/api/v3/ticker/price" 最新价格接口
    - "/api/v3/ticker/24hr"  24hr价格变动情况
    - "/api/v3/avgPrice" 当前平均价格
- trade_fee_service
- trade_service  
    - "/api/v3/historicalTrades" 历史成交
    - "/api/v3/aggTrades" 近期成交（归集）
    - "/api/v3/myTrades" 
- tradingDayTicker
    - "/api/v3/ticker/tradingDay"  # 交易日行情
- uiKlines_service
    - "/api/v3/uiKlines" UIK线数据
- user_stream_service
- user_universal_transfer
- websocket_service
- websocket
- withdraw_service
