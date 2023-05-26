#!/bin/bash

port=0
count="$2"
conc="$3"
ip="$4"

if [ "$1" == "trad" ]; then
    echo "testing traditional app..."
    port=8070
elif [ "$1" == "weaver" ]; then
    echo "testing weaver..."
    port=8000
else
    echo "please provide the system to test!"
    exit -1
fi

echo "benching $1 on $ip:$port with -n $count and -c $conc"
ab -n $count -c $conc http://"$ip":"$port"/request
