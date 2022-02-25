#!/usr/bin/env python3
"""mapper2.py"""

import sys
import string
import itertools

total = 0

for line in sys.stdin:
	line = line.strip()
	line_items = line.split(", ")
	line_items.sort()
	
	for i in range(0, len(line_items)-1):
		for j in range(i+1, len(line_items)):
			total += 1
			print('(%s, %s)\t%s' % (line_items[i], line_items[j], 1))

print('Total Unique Pairs:\t%s' % total)

	
	
