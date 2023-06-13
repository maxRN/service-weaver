#!/bin/bash

ip=$1
# weaver or trad
app=$2

# for n in 10000 50000 100000; do
#     for c in 1 10 20; do
        for ep in "long" "thick" "nothing"; do
            for test_run in 1 2 3; do
                echo "running bench $app $n $c $ip $ep"
                echo "test_run $test_run"
                ./bench.sh $app 10000 1 $ip $ep $test_run
                ./bench.sh $app 50000 10 $ip $ep $test_run
                ./bench.sh $app 100000 20 $ip $ep $test_run
            done
        done
#     done
# done
