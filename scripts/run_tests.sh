#!/bin/bash

ip=$1
# weaver or trad
app=$2

for n in 10000 50000 100000; do
    for c in 1 10 20; do
        for ep in "long" "thick" "nothing"; do
            for test_run in 1 2 3; do
                echo "running bench $app $n $c $ip $ep"
                echo "test_run $test_run"
                ./bench.sh $app $n $c $ip $ep
            done
        done
    done
done
