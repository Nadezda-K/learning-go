#!/bin/sh

file_counts=$(find . | wc -l)
multiplication=$(($file_counts * 5))

echo "\v\t Total files * 5: ${multiplication}\v\n"

