#! /bin/bash

# get the case number and save it to variable OUTPUT
OUTPUT=$(head -n 179 ./streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2) 


# echo the the case number
echo "$OUTPUT" 


# print out the content of the case we got from before
echo "$(cat ./interviews/interview-$OUTPUT)"


# filter out the vehicles by model, color and height of the person
grep -A 4 L337 ./vehicles | grep -A 3 -B 1 Honda | grep -A 2 -B 2 Blue | grep -B 4 "Height: 6"


# print out the matching number of memebrships according to the suspect
cat ./memberships/AAA ./memberships/Delta_SkyMiles ./memberships/Museum_of_Bash_History ./memberships/Terminal_City_Library| grep "$MAIN_SUSPECT" | wc -l
