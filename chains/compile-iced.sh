#!/bin/sh

args=`getopt s $*`
if [ $? -ne 0 ]
then
    usage
fi

PRETTY=json_pp

for i
do
	case "$i"
	in
		-s)
			PRETTY=cat
			shift
			;;
		--)
          shift ; break
			;;
	esac
done

for i in $*; do
	input=$i
	output=`basename -s .iced $i`.json
	forge-sigchain -f iced < $input | $PRETTY > $output
done

