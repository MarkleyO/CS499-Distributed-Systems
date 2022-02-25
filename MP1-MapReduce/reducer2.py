#!/usr/bin/env python3
"""recuder2.py"""

import sys

prev_set = None
current_count = 0
key = "Total Unique Pairs:"
unique = 0

for line in sys.stdin:
	line = line.strip()
	pair, count = line.split('\t', 1)
	count = int(count)

	#Since everything is coming in sorted, we can just leave the tuples 
	#as a string, and not worry about extracting individual data points

	if pair == key:
		unique += count

	if prev_set == pair:
		if prev_set != key:
			unique -= 1			
	
		current_count += int(count)
	else:
		if prev_set and prev_set != key:
			print('%s %s' % (prev_set, current_count))
		
		prev_set = pair
		current_count = count

if prev_set == pair:
	if prev_set != key:
		unique -= 1	
		print('%s %s' % (prev_set, current_count))

print('%s %s' % (key, unique))

