#!/bin/sh

for i in $*; do
	input=$i
	output=`basename -s .cson $i`.json
	../node_modules/.bin/forge-sigchain -f cson < $input | json_pp > $output
done

