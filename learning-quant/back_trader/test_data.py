# -*- coding:UTF-8 -*-
from __future__ import (absolute_import, division, print_function,
                        unicode_literals)

import pandas as pd
import backtrader as bt
import tushare as ts
from backtrader.feeds import GenericCSVData


class AddCsvData(GenericCSVData):
    lines = ('pe', 'gdp_growth', 'cpi',)
    params = (('pe', 7), ('gdp_growth', 8), ('cpi', 9),)


class TestStrategy1(bt.Strategy):
    def log(self, txt, dt=None):
        dt = dt or self.datas[0].datetime.date(0)
        print('%s, %s' % (dt.isoformat(), txt))

    def next(self):
        self.log(f"市盈率:{self.datas[0].pe[0]}, GDP:{self.datas[0].gdp_growth[0]}, CPI:{self.datas[0].cpi[0]}")


if __name__ == '__main__':
    # cerebro = bt.Cerebro()
    # cerebro.addstrategy(TestStrategy1)
    # feed = AddCsvData(dataname='test_idi.csv', dtformat=('%Y/%m/%d'))
    # cerebro.adddata(feed)
    # cerebro.run()
    df = pd.read_csv('test_idi.csv')
    df = df.sort_values('datetime', ascending=True)
    df.to_csv('filterOrder.csv', index=False)

