#!/bin/bash

ns=($((10**7)) $((10**8)) $((10**9)) )
ks=(1 2 4 8 16 32 64 128 256)

for n in "${ns[@]}"; do
    echo "++++++++++++++++++++++++++++++++++++++" >> go_results
    for k in "${ks[@]}"; do
	    go run ../main.go -n="$n" -k="$k" >> go_results
    done
done
