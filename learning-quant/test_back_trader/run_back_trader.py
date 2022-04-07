# import inline as inline
# import matplotlib
# import pandas as pd
# from datetime import datetime
# import backtrader as bt
# import tushare as ts
# import test_index_dynamic_invest as idi
#
#
# def get_data(code, start='2020-01-01', end='2020-03-31'):
#     df = ts.get_k_data(code, autype='qfq', start=start, end=end)
#     df.index = pd.to_datetime(df.date)
#     df['openinterest'] = 0
#     df = df[['open', 'high', 'low', 'close', 'volume', 'openinterest']]
#     return df
#
#
# if __name__ == '__main__':
#     dataframe = get_data('000905')
#
#     # 加载数据
#     raw_data = bt.feeds.PandasData(dataname=dataframe)
#
#     # 实例化 cerebro
#     cerebro = bt.Cerebro()
#     # 通过 feeds 读取数据
#     # data = btfeeds.BacktraderCSVData()
#     # 将数据传递给 “大脑”
#     cerebro.adddata(raw_data)
#     # 通过经纪商设置初始资金
#     cerebro.broker.setcash(100000)
#     # 设置单笔交易的数量
#     # cerebro.addsizer(1000)
#     # 设置交易佣金
#     cerebro.broker.setcommission(commission=0.02)
#     # 添加策略
#     cerebro.addstrategy(idi)
#     # 添加策略分析指标
#     # cerebro.addanalyzer()
#     # 添加观测器
#     # cerebro.addobserver()
#     # 启动回测
#     cerebro.run()
#     # 可视化回测结果
#     cerebro.plot()

# mpl.rcParams['axes.unicode_minus']=False

# 先引入后面可能用到的包（package）
import pandas as pd
from datetime import datetime
import backtrader as bt
import test_index_dynamic_invest as idi
# 使用tushare旧版接口获取数据
import tushare as ts


def get_data(code, start='2020-01-01', end='2020-03-31'):
    df = ts.get_k_data(code, autype='qfq', start=start, end=end)
    df.index = pd.to_datetime(df.date)
    df['openinterest'] = 0
    df = df[['open', 'high', 'low', 'close', 'volume', 'openinterest']]
    return df


dataframe = get_data('600000')

# 回测期间
start = datetime(2020, 1, 1)
end = datetime(2020, 3, 31)
# 加载数据
data = bt.feeds.PandasData(dataname=dataframe, fromdate=start, todate=end)

if __name__ == '__main__':
    # 初始化cerebro回测系统设置
    cerebro = bt.Cerebro()
    # 将数据传入回测系统
    cerebro.adddata(data)
    # 将交易策略加载到回测系统中
    cerebro.addstrategy(idi)
    # 设置初始资本为10,000
    startcash = 10000
    cerebro.broker.setcash(startcash)
    # 设置交易手续费为 0.2%
    cerebro.broker.setcommission(commission=0.002)

    d1 = start.strftime('%Y%m%d')
    d2 = end.strftime('%Y%m%d')
    print(f'初始资金: {startcash}\n回测期间：{d1}:{d2}')
    # 运行回测系统
    cerebro.run()
    # 获取回测结束后的总资金
    portvalue = cerebro.broker.getvalue()
    pnl = portvalue - startcash
    # 打印结果
    print(f'总资金: {round(portvalue, 2)}')
    print(f'净收益: {round(pnl, 2)}')
