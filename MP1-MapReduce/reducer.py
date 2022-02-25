#!/usr/bin/env python3
"""reducer.py"""

from operator import itemgetter # I guess this doesnt get used
import sys


current_word = None
current_count = 0
unique = 0
total = 0
unique_keyword = "Unique Words:"
total_keyword = "Total Words:"

# input comes from STDIN
for line in sys.stdin:
	# Clean the lines from whitespace
	line = line.strip()
	# take the word + tab + num format and break it up
	word, count = line.split('\t', 1)


	# reformat the key to an int
	try:
		count = int(count)
	except ValueError:
		# I don't think this should happen, but its protection
		# for is we get a non int
		continue
	
	if word == unique_keyword:
		unique += count
		
	if word == total_keyword:
		total += count

	# input should already be sorted by the key value (words)
	if current_word == word:
		current_count += count
		if current_word != unique_keyword and current_word != total_keyword:
			unique -= 1
			
	else:
		if current_word != unique_keyword and current_word != total_keyword:
			#results go to Standard Out
			print('%s %s' % (current_word, current_count))
		current_count = count
		current_word = word

if current_word == word and current_word != unique_keyword and current_word != total_keyword:
	if current_word != unique_keyword and current_word != total_keyword:
		unique -= 1
	print('%s %s' % (current_word, current_count))

print('%s %s' % ("Total Words:", total))
print('%s %s' % ("Unique Words:", unique))


