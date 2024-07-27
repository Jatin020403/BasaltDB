#!/bin/bash

# Record the start time
start_time=$(date +%s)
m=10000

# Loop to insert key-value pairs 100 times
for i in $(seq 1 $m)
do
  ./bin/basaltdb insert key$i value$i > /dev/null 2>&1
done

# Record the end time
end_time=$(date +%s)

# Calculate the time taken and print it
time_taken=$((end_time - start_time))
echo "Time taken for $m insertions: $time_taken seconds"


# Record the start time
start_time=$(date +%s)

# Loop to insert key-value pairs 100 times
for i in $(seq 1 $m)
do
  ./bin/basaltdb get key$i > /dev/null 2>&1
done

# Record the end time
end_time=$(date +%s)

# Calculate the time taken and print it
time_taken=$((end_time - start_time))
echo "Time taken for $m get: $time_taken seconds"
