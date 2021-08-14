#!/bin/bash
declare -a np_ns=(
    "1 1"
    "1 4"
    "1 8"
    "1 16"
    "2 1"
    "4 1"
    "8 1"
    "16 1"
)

N=(1 2 4 8 16 32)

N_LEN=${#N[@]}
NP_NC_LEN=${#NP_NC[@]}

for (( n=0; n<$N_LEN; n++)); 
do
    echo "++++++++++++++++++++++++++++++++++++++" >> go_results.txt
    for element in "${np_ns[@]}"; 
    do
        read -a np_nc <<< "$element"
        echo "Running for NP = ${np_nc[0]}; NC = ${np_nc[1]}; N= ${N[$n]}"

	    go run ../main.go -n=${N[$n]} -np=${np_nc[0]} -nc=${np_nc[1]} >> go_results.txt
    done
done
