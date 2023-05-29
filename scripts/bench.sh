#!/bin/bash

port=0
count="$2"
conc="$3"
ip="$4"
test_endpoint="$5"

if [ "$1" == "trad" ]; then
    echo "testing traditional app..."
    port=7766
elif [ "$1" == "weaver" ]; then
    echo "testing weaver..."
    port=8000
else
    echo "please provide the system to test!"
    exit -1
fi

csv_file_name="results/$1_$2_$3_$5.csv"
tsv_file_name="results/$1_$2_$3_$5.tsv"

echo "benching $1 on $ip:$port/$test_endpoint with -n $count and -c $conc"
if [ "$1" == "trad" ]; then
    ab -n $count -c $conc -e $csv_file_name -g $tsv_file_name http://"$ip":"$port"/"$test_endpoint"
else
    ab -n $count -c $conc -H "Host:hello.com" -e $csv_file_name -g $tsv_file_name http://"$ip":"$port"/"$test_endpoint"
fi
