from log.logger import get_logger
import app_context.AppContext
from app_context.AppContext import Context

logger = get_logger('index_dynamic_invest')


class IndexDynamicInvest:
    def __init__(self):
        """
        init
        """

    def indicators(self, context: Context):
        """
        指标
        :param context
        """
        context.trade_time_period = 7
        context.trade_amount = 1000
        context.last_trade_time = 1

    def choose_stock(self, context: Context):
        """
        标的,选择指数
        :param context
        """
        context.symbol_list = ['']

    def timing(self, context: Context):
        """
        择时
        :param context
        """
        # 1.买入
        # 1.1距离上次交易时间不超过T天,或者当前市盈率十年历史分位高于20%,不买入
        if context.last_trade_time <= context.trade_time_period or context.ttm_decade > 0.2:
            logger.info(
                f'buy strategy, last trade time:{context.last_trade_time} lq T:{context.trade_time_period} '
                f'or ttm十年历史分位:{context.ttm_decade} gt 20%')
            return
        # 1.2距离上次交易时间已超过T天
        if context.last_trade_time > context.trade_time_period:
            logger.info(f'buy strategy, last trade time:{context.last_trade_time} gt T:{context.trade_time_period}')
            # 1.2.1当前市盈率十年历史分位处于[10%, 20%],买入X金额
            if 0.1 <= context.ttm_decade <= 0.2:
                logger.info(f'buy strategy, ttm十年历史分位处于[10%, 20%], 买入金额:{context.trade_amount}')
                context.buy(context.trade_amount)
            # 1.2.2当前市盈率十年历史分位处于[0%, 10%],买入2X金额
            if 0 <= context.ttm_decade <= 0.1:
                logger.info(f'buy strategy, ttm十年历史分位处于[0%, 10%], 买入金额:2*{context.trade_amount}')
                context.buy(2 * context.trade_amount)
        # 2.卖出（针对指数基金）
        # 2.1距离上次交易时间不超过T天,或者当前市盈率十年历史分位低于50%,不卖出
        if context.last_trade_time <= context.trade_time_period or context.ttm_decade < 0.5:
            logger.info(
                f'sell strategy, last trade time:{context.last_trade_time} lq T:{context.trade_time_period} '
                f'or ttm十年历史分位:{context.ttm_decade} less 50% do not sell')
            return
        # 2.2距离上次交易时间已超过T天
        if context.last_trade_time > context.trade_time_period:
            logger.info(f'sell strategy, last trade time:{context.last_trade_time} gt T:{context.trade_time_period}')
            # 区间边界值？
            # 2.2.1当前市盈率十年历史分位处于[50%, 70%],卖出X金额
            if 0.5 <= context.ttm_decade <= 0.7:
                logger.info(f'sell strategy, ttm十年历史分位处于[50%, 70%], 卖出金额:{context.trade_amount}')
                context.sell(context.trade_amount)
            # 2.2.2当前市盈率十年历史分位处于[70%, 80%],卖出2X金额
            if 0.7 <= context.ttm_decade <= 0.8:
                logger.info(f'sell strategy, ttm十年历史分位处于[70%, 80%], 卖出金额:2*{context.trade_amount}')
                context.sell(2 * context.trade_amount)
            # 2.2.3当前市盈率十年历史分位处于80%以上区间,卖出所有金额
            if context.ttm_decade > 0.8:
                logger.info(f'sell strategy, ttm十年历史分位处于80%以上, 卖出所有金额')
                context.sell(-1)
        # 3.止损（无止损规则）
        # 4.止盈（等同于卖出规则）

    def control_risk(self, context: Context):
        """
        风控
        :param context
        """
        # 1.系统失效处理
        # 1.1当年GDP增长率小于当年CPI,不再增加买入
        if context.gdp_growth < context.cpi:
            logger.info(f'system invalid gdp growth:{context.gdp_growth} less cpi:{context.cpi}, not buy.')
            return False
        # 1.2连续两年GDP增长率小于CPI或者指数平均资本回报率低于GDP增长率,全部卖出
        if context.two_gdp_growth < context.two_cpi or context.roic < context.gdp_growth:
            logger.info(
                f'system invalid two year gdp growth:{context.two_gdp_growth} '
                f'less two year cpi:{context.two_cpi} sell all.')
            context.sell(-1)
            return False

    def run(self, context: Context):
        starter = IndexDynamicInvest()
        # indicators
        starter.indicators(context)
        # choose stock
        starter.choose_stock(context)
        # risk control
        if starter.control_risk(context):
            # timing
            starter.timing(context)
