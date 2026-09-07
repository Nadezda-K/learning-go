#!/bin/sh

touch a
touch \! 
touch \\ 
touch \"

mkdir \`

cp \! \`

echo $MOVE_A
if [ "$MOVE_A" = "yes" ]; then
    cp a \`
elif [ "$MOVE_A" = "no" ]; then
    rm a
else
   break
fi
