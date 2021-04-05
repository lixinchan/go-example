# -*- coding: utf-8 -*-
import math


class Func(object):

    @staticmethod
    def test_func():
        data = input('input number:')
        print(hex(int(data)))
        return None

    @staticmethod
    def test_none_func():
        # return
        # pass
        param = 3
        if isinstance(param, (int, float)):
            print('param type correct.')
        return 5,

    @staticmethod
    def test_func_param(x: int, n=2):
        pass

    @staticmethod
    def test_func_def_param(l=[]):
        l.append('end')
        return l

    @staticmethod
    def test_func_variable_param(*nums):
        ret = 0
        for idx in nums:
            ret += idx
        print(ret)

    @staticmethod
    def test_func_position_param(name, **kw):
        if 'test' in kw:
            pass
        print(name, kw)

    @staticmethod
    def test_func_named_param(name, *, test):
        print(name, test)

    @staticmethod
    def test_func_name_param(name, *args, test):
        print(name, args, test)


if __name__ == '__main__':
    # Func.test_func()
    # result = Func.test_none_func()
    # for idx in range(1, 3):
        # ret = Func.test_func_def_param()
        # print(idx, ret)
    # nums = (1, 2, 3, 4, 5)
    # Func.test_func_variable_param(*nums)
    # Func.test_func_position_param('test', test='test')
    # Func.test_func_named_param('test', test='test')
    Func.test_func_name_param('test', 1, 3, 5, test='test')
    pass
