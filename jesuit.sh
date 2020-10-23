for i in {5001..5010}
do
    go run node.go $i &
    sleep 1
done