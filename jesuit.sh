for i in {3001..}
do
    go run server.go $i &
    sleep 1
done