#!/bin/sh

loops_number=$1

if [[ $loops_number -gt 100 ]]; then
    $loops_number=100
fi

count_itr=1
while [ $count_itr -le $loops_number ]; do
    echo "This is loop number ${count_itr}"
    count_itr=$((count_itr + 1))
done
