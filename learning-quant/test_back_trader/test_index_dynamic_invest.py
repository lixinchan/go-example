from log.logger import get_logger
import backtrader as bt
import backtrader.indicators as btind
import backtrader.feeds as btfeeds

logger = get_logger('index_dynamic_invest')


class TestIndexDynamicInvest(bt.Strategy):
    params = (
        ('maperiod', 20),
    )

    def __init__(self):
        """
        指数动态投资法参数
        """
        # self.order = None
        # # 上次交易时间
        # self.last_trade_time = 6
        # # T天
        # self.trade_time_period = 7
        # # 市盈率十年历史分位
        # self.ttm_decade = 0.06
        # self.gdp_growth = 0.07
        # self.cpi = 0.056
        # self.two_gdp_growth = 0.075
        # self.two_cpi = 0.06
        # # 交易金额X
        # self.trade_amount = 1000

        # 指定价格序列
        self.dataclose = self.datas[0].close
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
            if self.dataclose[0] > self.sma[0]:
                # 执行买入
                self.order = self.buy(size=500)
        else:
            # 执行卖出条件判断：收盘价格跌破20日均线
            if self.dataclose[0] < self.sma[0]:
                # 执行卖出
                self.order = self.sell(size=500)

    # def next(self):
    #     """
    #     交易策略
    #     :return:
    #     """
    #     if self.order:
    #         return
    #     # 风控
    #     if not self.control_risk():
    #         return
    #         # 1.买入
    #     # 1.1距离上次交易时间不超过T天,或者当前市盈率十年历史分位高于20%,不买入
    #     if self.last_trade_time <= self.trade_time_period or self.ttm_decade > 0.2:
    #         logger.info(
    #             f'buy strategy, last trade time:{self.last_trade_time} lq T:{self.trade_time_period} '
    #             f'or ttm十年历史分位:{self.ttm_decade} gt 20%')
    #         self.order = None
    #     # 1.2距离上次交易时间已超过T天
    #     if self.last_trade_time > self.trade_time_period:
    #         logger.info(f'buy strategy, last trade time:{self.last_trade_time} gt T:{self.trade_time_period}')
    #         # 1.2.1当前市盈率十年历史分位处于[10%, 20%],买入X金额
    #         if 0.1 <= self.ttm_decade <= 0.2:
    #             logger.info(f'buy strategy, ttm十年历史分位处于[10%, 20%], 买入金额:{self.trade_amount}')
    #             self.order = self.buy(size=self.trade_amount)
    #         # 1.2.2当前市盈率十年历史分位处于[0%, 10%],买入2X金额
    #         if 0 <= self.ttm_decade <= 0.1:
    #             logger.info(f'buy strategy, ttm十年历史分位处于[0%, 10%], 买入金额:2*{self.trade_amount}')
    #             self.order = self.buy(size=2 * self.trade_amount)
    #     # 2.卖出（针对指数基金）
    #     # 2.1距离上次交易时间不超过T天,或者当前市盈率十年历史分位低于50%,不卖出
    #     if self.last_trade_time <= self.trade_time_period or self.ttm_decade < 0.5:
    #         logger.info(
    #             f'sell strategy, last trade time:{self.last_trade_time} lq T:{self.trade_time_period} '
    #             f'or ttm十年历史分位:{self.ttm_decade} less 50% do not sell')
    #         self.order = None
    #     # 2.2距离上次交易时间已超过T天
    #     if self.last_trade_time > self.trade_time_period:
    #         logger.info(f'sell strategy, last trade time:{self.last_trade_time} gt T:{self.trade_time_period}')
    #         # 区间边界值？
    #         # 2.2.1当前市盈率十年历史分位处于[50%, 70%],卖出X金额
    #         if 0.5 <= self.ttm_decade <= 0.7:
    #             logger.info(f'sell strategy, ttm十年历史分位处于[50%, 70%], 卖出金额:{self.trade_amount}')
    #             self.order = self.sell(size=self.trade_amount)
    #         # 2.2.2当前市盈率十年历史分位处于[70%, 80%],卖出2X金额
    #         if 0.7 <= self.ttm_decade <= 0.8:
    #             logger.info(f'sell strategy, ttm十年历史分位处于[70%, 80%], 卖出金额:2*{self.trade_amount}')
    #             self.order = self.sell(size=2 * self.trade_amount)
    #         # 2.2.3当前市盈率十年历史分位处于80%以上区间,卖出所有金额
    #         if self.ttm_decade > 0.8:
    #             logger.info(f'sell strategy, ttm十年历史分位处于80%以上, 卖出所有金额')
    #             self.order = self.sell(size=100000)
    #     # 3.止损（无止损规则）
    #     # 4.止盈（等同于卖出规则）
    #
    # def control_risk(self):
    #     """
    #     风控
    #     """
    #     # 1.系统失效处理
    #     # 1.1当年GDP增长率小于当年CPI,不再增加买入
    #     if self.gdp_growth < self.cpi:
    #         logger.info(f'system invalid gdp growth:{self.gdp_growth} less cpi:{self.cpi}, not buy.')
    #         self.order = None
    #         return False
    #     # 1.2连续两年GDP增长率小于CPI或者指数平均资本回报率低于GDP增长率,全部卖出
    #     if self.two_gdp_growth < self.two_cpi or self.roic < self.gdp_growth:
    #         logger.info(
    #             f'system invalid two year gdp growth:{self.two_gdp_growth} '
    #             f'less two year cpi:{self.two_cpi} sell all.')
    #         self.order = self.sell(size=100000)
    #         return False


