#!/bin/bash

# Find all PIDs associated with ./bombman
pids=$(pgrep -f "./bombman")

if [[ -z "$pids" ]]; then
  echo "No running bombman processes found."
  exit 1
fi

# Iterate over each PID and kill the process
for pid in $pids; do
  kill "$pid" && echo "Process with PID $pid has been killed."
done

echo "All bombman processes have been killed."