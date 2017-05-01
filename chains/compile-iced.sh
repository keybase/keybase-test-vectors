#!/bin/sh

for i in $*; do
	input=$i
	output=`basename -s .iced $1`.json
	forge-sigchain -f iced < $input | json_pp > $output
done

