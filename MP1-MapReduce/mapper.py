#!/usr/bin/env python3
"""mapper.py"""

import sys
import string

total = 0

# Input from Standard Input
for line in sys.stdin:
	# take away the whitespace on each line
	line = line.strip()
	line = line.lower()
	
	table = line.maketrans('', '',  string.punctuation)
	line = line.translate(table)

	# break line up into individual words
	words = line.split()

	for word in words:
		total += 1
		# the key for each word is just 1
		# this is then output to the reducer step
		print('%s\t%s' % (word, 1))

print('%s\t%s' % ("Total Words:", total))
print('%s\t%s' % ("Unique Words:", total))
