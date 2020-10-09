#!/bin/bash

URL="https://linuxnet.ca/ieee/oui/nmap-mac-prefixes"

TEMP_FILE=/tmp/nmap_oui.txt
TEMP_GOFILE=/tmp/oui_temp.go 
OUI_FILE=oui.go
wget -O $TEMP_FILE $URL -o /dev/null

echo 'package vendormap' > $TEMP_GOFILE
echo "" >> $TEMP_GOFILE
echo "// ManufacturerMap map with MAC vendor prefixes" >> $TEMP_GOFILE
echo 'var ManufacturerMap = map[string]string{' >> $TEMP_GOFILE

awk -F"	" '!_[$1]++' $TEMP_FILE >> $TEMP_FILE.tmp 
sed -e 's/\(.*\)	\(.*\)/"\1": "\2",/g' $TEMP_FILE.tmp >> $TEMP_GOFILE

echo  '}' >> $TEMP_GOFILE

rm $TEMP_FILE.tmp $TEMP_FILE
mv $TEMP_GOFILE $OUI_FILE
