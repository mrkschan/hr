#!/bin/bash

read -a INS <<< `find ins/ -type f | sort`
read -a OUTS <<< `find outs/ -type f | sort`

for idx in "${!INS[@]}"
do
	IN="${INS[idx]}"
	OUT="${OUTS[idx]}"

	cat $IN | python main.py > /tmp/out.txt
	DIFF=`diff $OUT /tmp/out.txt`

	if [ "$DIFF" != "" ]
	then
		echo "cat $IN | python3 main.py > /tmp/out.txt"
		echo "diff $OUT /tmp/out.txt"
		echo "$DIFF"
	fi
done
