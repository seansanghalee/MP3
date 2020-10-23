for i in {5001..5020}
do
    go run node.go $i &
    sleep .5
done