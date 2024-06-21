#!/bin/bash

go build

nohup ./bombman server > server.log 2>&1 &
server_pid=$!
echo "Server started with PID $server_pid"

# Wait for server to start
sleep 0.5

./bombman client > client1.log 2>&1 &
echo "Client 1 started with PID $!"

./bombman client > client2.log 2>&1 &
echo "Client 2 started with PID $!"

echo "All processes started and running indefinitely."
