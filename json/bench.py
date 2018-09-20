#!/usr/bin/env python -O
# -*- coding: utf-8 -*-
#from __future__ import print_function, unicode_literals, with_statement, division, absolute_import

import sys
import timeit
import json

import simplejson


def std_json(data: str):
    json.loads(data)


def simple_json(data: str):
    simplejson.loads(data)


def main():
    with open(sys.argv[1]) as fin:
        data = fin.read()

    benchmarks = [
        ('json.loads', std_json),
        ('simplejson.loads', simple_json),
    ]

    N = int(1e5)
    result = []
    for name, f in benchmarks:
        elapsed = timeit.timeit('f(data)', globals=locals(), number=N)
        result.append((name, elapsed))

    for name, elapsed in result:
        print(f'{name}: {elapsed}s')

    return 0


if __name__ == '__main__':
    sys.exit(main())
