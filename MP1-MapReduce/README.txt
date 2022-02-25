Version: Hadoop 3.2.2

First Part:
	Command: 
		mapred streaming -file mapper.py -mapper mapper.py \
		-file reducer.py -reducer reducer.py \
		-input /sample_input/* -output /sample_output/{name of output folder}
	
	Before running, use the hdfs dfs -copyFromLocal command to move the input .txt files
	into the directory /sample_input/. Also choose a unique name for the {name of output
	folder} field. This is where output will be stored.

Second Part:
	The only difference from the above command, should be replacing mapper and reducer with mapper 2 and reducer 2.
	Another file will also have to be input for this, formatted in the style described in the homework
	assignment document. 
