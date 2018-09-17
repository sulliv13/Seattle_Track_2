#!/bin/bash
FILES="AIS_2017_12_Zone17.zip
AIS_2017_12_Zone18.zip
AIS_2017_12_Zone11.zip
"
URL="https://coast.noaa.gov/htdata/CMSP/AISDataHandler/2017/"

for f in $FILES
do
if [ -e $f ]
then
    echo "$f already downloaded"
else
    echo "Downloading $URL$f..."
    wget $URL$f -q --show-progress     
fi
done

