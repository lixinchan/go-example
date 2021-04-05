#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import requests
import json
import re


class Basic(object):
    @staticmethod
    def test_print():
        print('hello', ',world', '!')
        print('100+200=', 100 + 200)
        print('1024', '*', '768', '=', 1024 * 768)
        print('''line1
        ...line2''')
        a = 'abc'
        b = a
        a = 'xyz'
        print(b)

    @staticmethod
    def test_data_type():
        print(ord('a'))
        print(ord('A'))
        print(chr(ord('A')))
        str_encode = '中文'.encode('utf-8')
        print(str_encode)
        # print('中文'.encode('ascii'))
        print(str_encode.decode('utf-8', errors='ignore'))
        print(len(b'abc'))
        print(len(str_encode))
        print(len('中文'))

        # format
        age = 'age: %s' % (28)
        print(age)

        print('{0}, {1}'.format(1, 2))

        s1 = 72
        s2 = 85
        r = (s2 - s1)/s1 * 100
        print('{0}, improve:{1:.1f}%'.format('xiaoming', r))

        print(f'{1}, {2}')

    @staticmethod
    def test_complex_type():
        list = ['1', '2']
        print(list)
        print(len(list))
        print(list[-1])
        list.append('3')
        print(list)
        idx = list.pop(1)
        print(idx)
        list.append(4)
        print(list)

        classmates = ('xiaoming', 'xiaoli')
        print(classmates)
        print(len(classmates))
        for e in classmates:
            print(e)
        t = ()
        print(t)
        t = (1)
        print(t)
        t = (1,)
        print(t)
        t = ('a', 'b', ['A', 'B'])
        print(t)
        t[2][0] = 'X'
        t[2][1] = 'Y'
        print(t)

    @staticmethod
    def test_regex():
        regex = '[A-Za-z0-9]+[.][com|cn|com.cn]+'
        target = 'safafsaf.jd.com'
        ret = re.search(regex, target)
        print(ret.start())
        print(ret.end())
        print(target[ret.start():ret.end()])

    @staticmethod
    def test_condition():
        age = input('age:')
        age = int(age)
        if age > 18:
            print('young')
        elif age > 18 and age < 25:
            print('old')
        else:
            print('older')

    @staticmethod
    def test_loop():
        sum = 0
        for idx in range(101):
            sum += idx
        print('sum', sum)

        L = ['Bart', 'Lisa', 'Adam']
        for v in L:
            print(f'Hello, {v}!')
        return None

    @staticmethod
    def test_dict():
        d = {1: 1, 2: 2, 3: 3}
        print(type(d))
        l = [1, 2, 3]
        print(type(set(l)))

    @staticmethod
    def test_list_map():
        test_map = {'test': 'test-1', 'a': 'a'}
        test_list = [test_map]
        ret = [dict(ele).get('test') for ele in test_list]
        print(ret[0])

        result = [dict(ele).get('a') for ele in test_list]
        print(result[0])


if __name__ == '__main__':
    # Basic.test_print()
    # Basic.test_data_type()
    # Basic.test_complex_type()
    # Basic.test_regex()
    # Basic.test_condition()
    # Basic.test_loop()
    # Basic.test_dict()
    Basic.test_list_map()
