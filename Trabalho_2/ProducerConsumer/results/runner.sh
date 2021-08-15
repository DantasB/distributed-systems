#!/bin/bash
declare -a np_ns=(
    "1 1"
    "1 2"
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

mkdir -p runs
LAST_ARCHIVE_DIR=runs

echo "time(s),n,np,nc" > go_results.csv
for (( n=0; n<$N_LEN; n++)); 
do
    for element in "${np_ns[@]}"; 
    do
        read -a np_nc <<< "$element"
        echo "Running for NP = ${np_nc[0]}; NC = ${np_nc[1]}; N = ${N[$n]}"
        for (( i=0; i<10;i++));
        do
            current_archive="$LAST_ARCHIVE_DIR/NP=${np_nc[0]};NC=${np_nc[1]};N=${N[$n]};RUN=$i.txt"
	        go run ../main.go -n=${N[$n]} -np=${np_nc[0]} -nc=${np_nc[1]} > $current_archive
            echo "$(tail -n 1 $current_archive),${N[$n]},${np_nc[0]},${np_nc[1]}" >> go_results.csv
        done
    done
done
