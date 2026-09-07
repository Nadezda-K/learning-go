#!/bin/sh

file_counts=$(find . | wc -l)
multiplication=$(($file_counts * 5))

printf "\t\vTotal files * 5: %d\v\n" "${multiplication}"

