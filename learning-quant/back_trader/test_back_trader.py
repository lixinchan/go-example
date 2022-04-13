# -*- coding:UTF-8 -*-
from __future__ import (absolute_import, division, print_function,
                        unicode_literals)

import backtrader as bt
import pandas as pd
import tushare as ts
from datetime import datetime


class TestIndexDynamicInvest(bt.Strategy):
    params = (
        ('maperiod', 20),
    )

    def __init__(self):
        """
        单均线策略参数
        """
        # 指定价格序列
        self.data_close = self.datas[0].close
        # 初始化交易指令、买卖价格和手续费
        self.order = None
        self.buyprice = None
        self.buycomm = None

        # 添加移动均线指标，内置了talib模块
        self.sma = bt.indicators.SimpleMovingAverage(
            self.datas[0], period=self.params.maperiod)

    def next(self):
        if self.order:  # 检查是否有指令等待执行,
            return
            # 检查是否持仓
        if not self.position:  # 没有持仓
            # 执行买入条件判断：收盘价格上涨突破20日均线
            if self.data_close[0] > self.sma[0]:
                # 执行买入
                self.order = self.buy(size=500)
        else:
            # 执行卖出条件判断：收盘价格跌破20日均线
            if self.data_close[0] < self.sma[0]:
                # 执行卖出
                self.order = self.sell(size=500)


def get_data(code, start='2021-01-01', end='2022-04-07'):
    df = ts.get_k_data(code, start=start, end=end, index=True)
    df.index = pd.to_datetime(df.date)
    df['openinterest'] = 0
    df = df[['open', 'high', 'low', 'close', 'volume', 'openinterest']]
    return df

    df['openinterest'] = 0
    df = df[['trade_time', 'open', 'high', 'low', 'close', 'volume', 'openinterest']]
    return df


dataframe = get_data('000905')

# 回测期间
start = datetime(2021, 1, 1)
end = datetime(2022, 4, 7)
# 加载数据
data = bt.feeds.PandasData(dataname=dataframe, fromdate=start, todate=end)

# if __name__ == '__main__':
#     # 初始化cerebro回测系统设置
#     cerebro = bt.Cerebro()
#     # 将数据传入回测系统
#     cerebro.adddata(data)
#     # 将交易策略加载到回测系统中
#     cerebro.addstrategy(TestIndexDynamicInvest)
#     # 初始资本为10,000
#     start_cash = 10000
#     cerebro.broker.setcash(start_cash)
#     # 交易手续费
#     cerebro.broker.setcommission(commission=0.02)
#
#     d1 = start.strftime('%Y%m%d')
#     d2 = end.strftime('%Y%m%d')
#     print('期初资金:%d,回测期间:%s-%s' % (start_cash, d1, d2))
#     # 运行回测系统
#     cerebro.run()
#     # 获取回测结束后的总资金
#     port_value = cerebro.broker.getvalue()
#     pnl = port_value - start_cash
#     # 打印结果
#     print('期末资金:%.2f' % port_value)
#     print('净收益:%.2f' % pnl)
if __name__ == '__main__':
    get_data('000905').to_csv('test.csv', index=False)
