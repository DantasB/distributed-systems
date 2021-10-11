
COORDINATOR_PATH=Coordinator/
PROCESS_PATH=Process/
VALIDATION_PATH=Validation/
cd $PROCESS_PATH
 > resultado.txt

NUMBER_OF_PROCESSES=$1
R=$2
K=$3
for PROCESS_NUMBER in $(seq 1 $NUMBER_OF_PROCESSES);
do
    go run process.go -pn=$PROCESS_NUMBER -r=$R -k=$K &
done

wait

cd ../$VALIDATION_PATH
go run validator.go -r=$R -n=$NUMBER_OF_PROCESSES