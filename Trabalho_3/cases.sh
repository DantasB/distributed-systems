COORDINATOR_PATH=Coordinator/
VALIDATION_PATH=Validation/
PROCESS_PATH=Process/
RESULTS_PATH=Results
cd $PROCESS_PATH/

mkdir -p $RESULTS_PATH/

R=5
K=1
NS=(2 4 8 16 32 64)
for N in ${NS[@]};
do
     > resultado.txt
    for PROCESS_NUMBER in $(seq 1 $N);
    do
        go run process.go -pn=$PROCESS_NUMBER -r=$R -k=$K &
    done
    sleep $(($K*$R*$N*2))
    cp resultado.txt $RESULTS_PATH/resultado_$N.txt
    cd ../$VALIDATION_PATH
    go run validator.go -r=$R -n=$N
    cd ../$PROCESS_PATH
done