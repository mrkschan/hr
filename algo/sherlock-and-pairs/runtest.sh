#!/bin/bash

read -a INS <<< `find ins/ -type f | sort`
read -a OUTS <<< `find outs/ -type f | sort`

EXEC="go run main.go"
# EXEC="python main.py"

for idx in "${!INS[@]}"
do
	IN="${INS[idx]}"
	OUT="${OUTS[idx]}"

	cat $IN | $EXEC > /tmp/out.txt
	DIFF=`diff $OUT /tmp/out.txt`

	if [ "$DIFF" != "" ]
	then
		echo "cat $IN | $EXEC > /tmp/out.txt"
		echo "diff $OUT /tmp/out.txt"
		echo "$DIFF"
	fi
done
