#!/bin/sh

input=$1
output=`basename -s .iced $1`.json
forge-sigchain -f iced < $input | json_pp > $output

