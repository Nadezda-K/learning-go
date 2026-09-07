#!/bin/sh

file_counts=$(find . | wc -l)
multiplication=$(($file_counts * 5))

echo $file_counts $multiplication
printf "\v\t Total files * 5: %d\v\n" "${multiplication}"

